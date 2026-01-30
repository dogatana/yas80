package binwriter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"yas80/errcode"
	"yas80/logging"
	"yas80/object"
)

type Segment struct {
	allocType int
	addr      int
	size      int // children も含むサイズ
	code      []byte
	children  []*Segment
}

func (s *Segment) String() string {
	return fmt.Sprintf("Segment $%04x - $%04x, size: $%x, children[%d]",
		s.addr, s.addr+s.size-1, s.size, len(s.children))
}

type BinWriter struct {
	prog   *object.BlockObject
	logger *logging.Logger
}

func New(prog object.Object, logger *logging.Logger) *BinWriter {
	return &BinWriter{prog: prog.(*object.BlockObject), logger: logger}
}

func (b *BinWriter) Write(filename string) error {
	segs, err := b.allocateMemory()
	if err != nil {
		fmt.Printf("error %s", err.Error())
	}
	for i, s := range segs {
		fmt.Printf("%d: %s\n", i, s.String())
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	var size, n int
	w := bufio.NewWriter(f)
	for _, s := range segs {
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
	w.Flush()
	fmt.Printf("write %d(0x%x)\n", size, size)
	return nil
}

func (b *BinWriter) allocateMemory() ([]*Segment, error) {
	segs := b.collectSegemnts()
	nsegs := make([]*Segment, 0, len(segs))

	if len(segs) == 0 {
		return nsegs, nil
	}

	if segs[0].allocType != 0 {
		return nsegs, errors.New(errcode.EBW_NOT_ABS_FIRST)
	}

	seg := segs[0]
	for _, s := range segs[1:] {
		if s.allocType != 0 {
			seg.children = append(seg.children, s)
			seg.size += s.size
			continue
		}

		naddr := seg.addr + seg.size
		switch {
		case s.addr < naddr:
			return nsegs, fmt.Errorf(errcode.EBW_OVERLAPPED, seg.addr, s.addr)
		case s.addr > naddr:
			size := s.addr - naddr
			code := make([]byte, size)
			for i := 0; i < size; i++ {
				code[i] = 0 // TODO: fill_byte で埋める
			}
			fill := &Segment{addr: naddr, code: code, size: size}
			nsegs = append(nsegs, seg, fill)
			seg = s
		default:
			nsegs = append(nsegs, seg)
			seg = s
		}
	}
	nsegs = append(nsegs, seg)
	return nsegs, nil
}

func (b *BinWriter) collectSegemnts() []*Segment {
	objs := b.flattenObject(b.prog)
	inseg := false
	segs := []*Segment{}

	var seg *Segment
	for _, o := range objs {
		switch o := o.(type) {
		case *object.CodeObject:
			if inseg {
				// セグメント処理中
				seg.code = append(seg.code, o.Code...)
				seg.size += len(o.Code)
				continue
			}
			// ORG なしに CODE 開始した場合、新規にセグメント作成
			seg = &Segment{addr: o.Addr, code: slices.Clone(o.Code), size: len(o.Code)}
			inseg = true
		case *object.OrgObject:
			// 処理中のセグメントを保存
			if seg != nil {
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
