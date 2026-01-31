package binwriter

import (
	"errors"
	"fmt"
	"io"
	"slices"
	"yas80/errcode"
	"yas80/object"
)

type Segment struct {
	allocType int
	addr      int
	size      int // children も含むサイズ
	code      []byte
	gap       bool // gap のために作成された Segment
	children  []*Segment
	err       error
}

func (s *Segment) String() string {
	if s.gap {
		return fmt.Sprintf("GAP %04x - %04x, size: $%x", s.addr, s.addr+s.size-1, s.size)

	}
	return fmt.Sprintf("Segment $%04x - $%04x, size: $%x, children[%d]",
		s.addr, s.addr+s.size-1, s.size, len(s.children))
}

type BinWriter struct {
	prog *object.BlockObject
	segs []*Segment // 配置済み Segment
}

func New(prog object.Object) *BinWriter {
	return &BinWriter{prog: prog.(*object.BlockObject)}
}

func (b *BinWriter) Write(w io.Writer) error {
	if b.segs == nil {
		return errors.New(errcode.EBW_NULL)
	}
	return b.write(w)
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
func (b *BinWriter) Allocate() error {
	segs := b.collectSegemnts()
	newSegs := make([]*Segment, 0, len(segs))

	if len(segs) == 0 {
		return nil
	}

	var seg *Segment
	if segs[0].allocType != 0 {
		// 最初が REL Segment の場合、addr:0 で Segment を作成し、children に設定
		seg = &Segment{addr: 0, size: segs[0].size, children: []*Segment{segs[0]}}
	} else {
		seg = segs[0]
	}

	// 以降のセグメントを割り当て
	for _, s := range segs[1:] {
		if s.allocType != 0 {
			// REL
			seg.children = append(seg.children, s)
			seg.size += s.size
			continue
		}

		naddr := seg.addr + seg.size // 現在の Sgement の次の開始アドレス
		switch {
		// ABS Segment のアドレス重複
		case s.addr < naddr:
			newSegs = append(newSegs, seg, s)
			b.segs = newSegs
			return fmt.Errorf(errcode.EBW_OVERLAPPED, seg.addr, s.addr)

		// ABS Segment 間にギャップあり
		case s.addr > naddr:
			size := s.addr - naddr
			code := make([]byte, size)
			for i := 0; i < size; i++ {
				code[i] = 0 // TODO: fill_byte で埋める
			}
			fill := &Segment{addr: naddr, code: code, size: size, gap: true}
			newSegs = append(newSegs, seg, fill)
			seg = s

		default:
			newSegs = append(newSegs, seg)
			seg = s
		}
	}
	newSegs = append(newSegs, seg)

	b.segs = newSegs
	return nil
}

// OrgObject, CodeObject から Segment を作成
func (b *BinWriter) collectSegemnts() []*Segment {
	objs := b.flattenObject(b.prog)
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

// prog(*object.BlockObject) を単一階層の slice に変換
func (b *BinWriter) flattenObject(obj object.Object) []object.Object {
	objs := []object.Object{}

	switch obj := obj.(type) {
	case *object.BlockObject:
		for _, o := range obj.Block {
			objs = append(objs, b.flattenObject(o)...)
		}
	default:
		objs = append(objs, obj)
	}
	return objs
}
