package filecontent

import "fmt"

type Context struct {
	FileContent *FileContent
	Source      *Context // マクロ定義 Context
	Line        uint32   // 1-
	Index       uint32   // 0- FileBox.Contet の Index
	Offset      uint32   // 0 マクロ未展開, 1- マクロ展開後
}

func (c *Context) String() string {
	if c == nil {
		return "  :  (  )"
	}
	ret := fmt.Sprintf("%2d:%2d", c.Line, c.Offset)
	if c.Source == nil {
		ret += "(  )"
	} else {
		ret += fmt.Sprintf("(%2d)", c.Source.Line)
	}
	return ret
}

func (c *Context) Equal(o *Context) bool {
	if o == nil {
		return false
	}
	return c.FileContent == o.FileContent && c.Line == o.Line && c.Index == o.Index && c.Offset == o.Offset
}

// Context の示す行を取得
func (c *Context) GetLine() string {
	if c == nil {
		return "<Context nil>"
	}
	if c.Offset == 0 {
		s, _ := c.FileContent.GetLine(int(c.Line))
		return s
	}
	if c.Source != nil {
		s, _ := c.Source.FileContent.GetLine(int(c.Source.Line))
		return s
	}
	return "<Context.Source nil>"
}
