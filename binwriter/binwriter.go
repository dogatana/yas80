package binwriter

import (
	"bytes"
	"fmt"
	"io"
	"slices"
	"strings"
	"yas80/errcode"
	"yas80/internal/util"
	"yas80/logging"
	"yas80/object"
)

type Segment struct {
	allocType int
	addr      int
	size      int // children も含むサイズ
	code      []byte
	gap       bool // gap のために作成された Segment
	children  []*Segment
}

func (s *Segment) String() string {
	if s.gap {
		return fmt.Sprintf("GAP %04x - %04x, size: $%x", s.addr, s.addr+s.size-1, s.size)

	}
	return fmt.Sprintf("Segment $%04x - $%04x, size: $%x, children[%d]",
		s.addr, s.addr+s.size-1, s.size, len(s.children))
}

type BinWriter struct {
	prog   *object.BlockObject
	segs   []*Segment // 配置済み Segment
	fill   int
	start  int // 実行開始アドレス
	logger *logging.Logger
}

func New(prog object.Object, fill int, logger *logging.Logger) *BinWriter {
	return &BinWriter{prog: prog.(*object.BlockObject), fill: fill, start: -1, logger: logger}
}

// BIN 形式の出力
func (b *BinWriter) WriteBin(w io.Writer) bool {
	b.allocateSegments()
	if b.logger.ErrorCount() > 0 {
		return false
	}
	if len(b.segs) == 0 {
		b.logger.Error(errcode.EBW_NULL, nil)
		return false
	}
	b.write(w)
	return true
}

// T88 形式の出力
func (b *BinWriter) WriteT88(w io.Writer, name string) bool {
	var buf bytes.Buffer
	if !b.WriteBin(&buf) {
		return false
	}
	data := buf.Bytes()

	t88size := 24 + 6 + 12*3 + 25 + 12                        // magic + VER/BLANK/SPACE/MARK/DATA(name)/MARK
	t88size += 16 + 4 + (len(data)+254)/255*3 + len(data) + 3 // start addr + data w/metadata + end
	t88size += 12*3 + 4                                       // MARK/SPACE/BLANK/ENd
	t88 := make([]byte, 0, t88size)

	// magic
	t88 = append(t88, []byte("PC-8801 Tape Image(T88)\x00")...)

	// VER
	t88 = append(t88, []byte{0x01, 0x00, 0x02, 0x00, 0x00, 0x01}...)

	// BLANK
	tick := 0
	tksize := 0x01e0
	t88 = append(t88, t88FixedSizeTag(0x0100, tick, tksize)...)
	tick += tksize

	// SPACE
	tksize = 0x12c0
	t88 = append(t88, t88FixedSizeTag(0x0102, tick, tksize)...)
	tick += tksize

	// MARK
	tksize = 0x2ee0
	t88 = append(t88, t88FixedSizeTag(0x0103, tick, tksize)...)
	tick += tksize

	// DATA - load name
	var ln string
	if !util.IsAsciiString(name) { // アスキー文字のみ有効項
		b.logger.Warning(errcode.WBW_LOAD_NAME, nil)
		ln = "$$$T88BIN"
	} else {
		ln = ("$$$" + name + "      ")[:9] // 6 文字のみ有効。足りない場合は空白で埋める
	}
	tksize = 9 * 44
	t88 = append(t88, 0x01, 0x01, 0x15, 0x00)   // id, size
	t88 = append(t88, intToBytes(tick, 4)...)   // tick
	t88 = append(t88, intToBytes(tksize, 4)...) // tick length
	t88 = append(t88, 0x09, 0x00, 0xcc, 0x01)   // data size, attr
	t88 = append(t88, []byte(ln)...)            // load name
	tick += tksize

	// MARK
	tksize = 0x03c0
	t88 = append(t88, t88FixedSizeTag(0x0103, tick, tksize)...)
	tick += tksize

	// DATA - binary
	addr := b.segs[0].addr // load addr
	tag := t88DataTag(data, tick, addr)
	t88 = append(t88, tag...)
	tick += int(tag[8]) + int(tag[9])*256 // 生成済みの tag length

	// MARK
	tksize = 0x1860
	t88 = append(t88, t88FixedSizeTag(0x0103, tick, tksize)...)
	tick += tksize

	// SPACE
	tksize = 0x2ee0
	t88 = append(t88, t88FixedSizeTag(0x0102, tick, tksize)...)
	tick += tksize

	// BLANK
	tksize = 0x3fc0
	t88 = append(t88, t88FixedSizeTag(0x0100, tick, tksize)...)

	// END
	t88 = append(t88, 0, 0, 0, 0)

	if s, err := w.Write(t88); err != nil || s != len(t88) {
		b.logger.Error(fmt.Sprintf(errcode.EBW_WRITE, err.Error()), nil)
		return false
	}
	return true
}

// 固定長 tag BLANK/SPACE/MARK を生成
func t88FixedSizeTag(tag, tick, tksize int) []byte {
	out := []byte{byte(tag & 0xff), byte(tag >> 8), 0x08, 0x00} // tag size == 8
	out = append(out, intToBytes(tick, 4)...)
	out = append(out, intToBytes(tksize, 4)...)
	return out
}

// DATA tag を生成
func t88DataTag(data []byte, tick, addr int) []byte {
	recSize := 7 + (len(data)+254)/255*3 + len(data) // ':' 以降の record size
	tagSize := recSize + 12

	out := make([]byte, 0, recSize+16)
	out = append(out, 0x01, 0x01, byte(tagSize&0xff), byte((tagSize>>8)&0xff)) // id, size
	out = append(out, intToBytes(tick, 4)...)                                  // tick
	out = append(out, intToBytes(recSize*44, 4)...)                            // tick length
	out = append(out, byte(recSize&0xff), byte((recSize>>8)&0xff), 0xcc, 0x01) // data size, attr
	slow := byte(addr & 0xff)
	shigh := byte((addr >> 8) & 0xff)
	out = append(out, ':', slow, shigh, slow+shigh) // start addr
	for ofs := 0; ofs < len(data); ofs += 255 {
		out = append(out, bytesToRecord(data[ofs:min(ofs+255, len(data))])...)
	}
	out = append(out, ':', 0, 0)
	return out
}

// 最大 255 の []byte を 3a レコードに変換する
func bytesToRecord(data []byte) []byte {
	if len(data) > 255 {
		panic(fmt.Sprintf("exceeds max data length 0x%x", len(data)))
	}
	out := make([]byte, len(data)+3)
	out[0] = ':'
	out[1] = byte(len(data))
	copy(out[2:], data)
	csum := byteSum(out[1 : len(out)-1])
	out[len(out)-1] = csum
	return out
}

// []byte の合計を byte として返す
func byteSum(data []byte) byte {
	var total byte
	for _, b := range data {
		total += b
	}
	return total
}

// tick, tick length -> []byte 変換ヘルパ関数
func intToBytes(data, size int) []byte {
	out := make([]byte, size)
	for i := range size {
		out[i] = byte(data & 0xff)
		data >>= 8
	}
	return out
}

// MZT 形式の出力
func (b *BinWriter) WriteMzt(w io.Writer, name string, start int) bool {
	var buf bytes.Buffer
	if !b.WriteBin(&buf) {
		return false
	}
	data := buf.Bytes()

	// header
	header := make([]byte, 128)

	// mode
	header[0] = 0x01 // 固定

	// load name
	// アスキー文字のみ有効項
	if !util.IsAsciiString(name) {
		b.logger.Warning(errcode.WBW_LOAD_NAME, nil)
		name = "OUTPUT"
	}
	name += "\r" + strings.Repeat(" ", 15) // 16文字に足りない場合空白で埋める
	copy(header[1:], []byte(name[:16]))
	header[0x11] = 0x0d

	// binary size
	size := len(data)
	header[0x12] = byte(size & 0xff)
	header[0x13] = byte((size >> 8) & 0xff)

	// load addr
	load := b.segs[0].addr
	header[0x14] = byte(load & 0xff)
	header[0x15] = byte((load >> 8) & 0xff)

	// start addr
	switch {
	case start == -1 && b.start == -1: // オプション指定なし、END 指定なし
		start = load
	case start == -1 && b.start != -1: // オプション指定なし、END 指定あり
		start = b.start
	case start != -1: // オプション指定あり
		// do nothing
	default:
		panic("cannot define start")
	}
	header[0x16] = byte(start & 0xff)
	header[0x17] = byte((start >> 8) & 0xff)

	// mzt header
	if s, err := w.Write(header); err != nil || s != len(header) {
		b.logger.Error(fmt.Sprintf(errcode.EBW_WRITE, err.Error()), nil)
		return false
	}

	// binary data
	if s, err := w.Write(data); err != nil || s != len(data) {
		b.logger.Error(fmt.Sprintf(errcode.EBW_WRITE, err.Error()), nil)
		return false
	}
	return true
}

func (b *BinWriter) WriteMap(w io.Writer) error {
	ofs := 0
	for _, s := range b.segs {
		if s.gap {
			if _, err := fmt.Fprintf(w, "%05x GAP %04x - %04x, size %04x\n", ofs, s.addr, s.addr+s.size-1, s.size); err != nil {
				return err
			}
			ofs += s.size
		} else {
			size := len(s.code)
			if _, err := fmt.Fprintf(w, "%05x ABS %04x - %04x, size %04x\n", ofs, s.addr, s.addr+size-1, size); err != nil {
				return err
			}
			ofs += size
			for _, rs := range s.children {
				if _, err := fmt.Fprintf(w, "%05x REL %04x - %04x, size %04x\n", ofs, rs.addr, rs.addr+rs.size-1, size); err != nil {
					return err
				}
				ofs += rs.size
			}
		}
	}
	return nil
}

func (b *BinWriter) write(w io.Writer) error {

	var (
		size, n int
		err     error
	)

	for _, s := range b.segs {
		if n, err = w.Write(s.code); err != nil {
			return err
		}
		size += n
		for _, cs := range s.children {
			if n, err = w.Write(cs.code); err != nil {
				return err
			}
			size += n
		}
	}
	return nil
}

// Segment を割り付ける
func (b *BinWriter) allocateSegments() {
	segs := b.collectSegemnts()

	// 有効 Segment なし
	if len(segs) == 0 {
		b.segs = segs
		return
	}
	segs = b.mergeREL(segs)

	// addr 順にソート
	slices.SortFunc(segs, func(a, b *Segment) int { return a.addr - b.addr })

	nsegs := make([]*Segment, 1, len(segs))
	nsegs[0] = segs[0]

	addr := segs[0].addr + segs[0].size
	for i, s := range segs[1:] {
		switch {
		case addr < s.addr:
			size := s.addr - addr
			code := make([]byte, size)
			for i := 0; i < size; i++ {
				code[i] = byte(b.fill)
			}
			gap := &Segment{addr: addr, code: code, size: size, gap: true}
			nsegs = append(nsegs, gap, s)
			addr += size + s.size

		case addr > s.addr:
			nsegs = append(nsegs, s)
			b.segs = nsegs
			b.logger.Error(fmt.Sprintf(errcode.EBW_OVERLAPPED, segs[i].addr, s.addr), nil)
			if addr < s.addr+s.size {
				// 現在の Segment の後続アドレスを設定する
				addr = s.addr + s.size
			}
		default:
			nsegs = append(nsegs, s)
			addr += s.size
		}
	}

	b.segs = nsegs
}

// REL を ABS に merge
func (b *BinWriter) mergeREL(segs []*Segment) []*Segment {
	nsegs := make([]*Segment, 0, len(segs))

	var seg *Segment

	if segs[0].allocType != 0 {
		// 最初が REL Segment の場合、addr:0 で Segment を作成し、children に設定
		seg = &Segment{addr: 0, size: segs[0].size, children: []*Segment{segs[0]}}
	} else {
		seg = segs[0]
	}

	for _, s := range segs[1:] {
		if s.allocType != 0 { // REL
			seg.children = append(seg.children, s)
			seg.size += s.size
		} else {
			nsegs = append(nsegs, seg)
			seg = s
		}
	}
	nsegs = append(nsegs, seg)

	return nsegs
}

// OrgObject, CodeObject から Segment を作成
func (b *BinWriter) collectSegemnts() []*Segment {
	objs := object.FlattenObject(b.prog)
	inseg := false // Segment 処理中
	segs := []*Segment{}

	var seg *Segment // 処理中 Segment
	for _, o := range objs {
		switch o := o.(type) {
		case *object.CodeObject:
			if inseg {
				// 処理中の Segment に追加
				seg.code = append(seg.code, o.Code...)
				seg.size += len(o.Code)
				continue
			}
			// 最初が CodeObject の場合、新規に Segment を作成
			seg = &Segment{addr: o.Addr, code: slices.Clone(o.Code), size: len(o.Code)}
			inseg = true
		case *object.OrgObject:
			if seg != nil {
				// 処理中の Segment があれば保存
				segs = append(segs, seg)
			}
			seg = &Segment{addr: o.Addr, code: []byte{}, allocType: o.AllocType}
			inseg = true
		case *object.EndObject:
			if o.Start != -1 {
				b.start = o.Start // 開始アドレスの更新
			}
		}
	}
	if inseg {
		segs = append(segs, seg)
	}
	return segs
}
