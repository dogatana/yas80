package fileblock

import "fmt"

type Context struct {
	FileBlock *FileBlock
	Line      int      // 1-
	Index     int      // 0- FileBox.Contet の Index
	Offset    int      // 0 マクロ未展開, 1- マクロ展開後
	Source    *Context // マクロ定義 Context
}

func (c *Context) String() string {
	ret := fmt.Sprintf("%q:%d,%d", c.FileBlock.Filename, c.Line, c.Offset)
	if c.Source != nil {
		ret += fmt.Sprintf("[%q:%d,%d)]", c.Source.FileBlock.Filename, c.Source.Line, c.Source.Offset)
	}
	return ret
}

func (c *Context) Equal(o *Context) bool {
	return c.FileBlock == o.FileBlock && c.Line == o.Line && c.Index == o.Index && c.Offset == o.Offset
}
