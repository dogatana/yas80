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
	logger *logging.Logger
}

func New(prog object.Object, fill int, logger *logging.Logger) *BinWriter {
	return &BinWriter{prog: prog.(*object.BlockObject), fill: fill, logger: logger}
}

// BIN 形式の出力
func (b *BinWriter) WriteBin(w io.Writer) bool {
	b.allocate()
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
	header[0] = 0x01
	// load name
	// アスキー文字のみ有効項
	if !util.IsAsciiString(name) {
		b.logger.Warning(errcode.WBW_LOAD_NAME, nil)
		name = "OUTPUT"
	}
	name += strings.Repeat(" ", 15) // 16文字に足りない場合空白で埋める
	copy(header[1:], []byte(name[:16]))
	header[0x11] = 0x0d
	// size
	size := len(data)
	header[0x12] = byte(size & 0xff)
	header[0x13] = byte((size >> 8) & 0xff)
	// load
	load := b.segs[0].addr
	header[0x14] = byte(load & 0xff)
	header[0x15] = byte((load >> 8) & 0xff)
	// start
	if start < 0 {
		start = load
	}
	header[0x16] = byte(start & 0xff)
	header[0x17] = byte((start >> 8) & 0xff)

	if s, err := w.Write(header); err != nil || s != len(header) {
		b.logger.Error(fmt.Sprintf(errcode.EBW_WRITE, err.Error()), nil)
		return false
	}
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
func (b *BinWriter) allocate() {
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
		}
	}
	if inseg {
		segs = append(segs, seg)
	}
	return segs
}
