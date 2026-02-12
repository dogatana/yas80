package filecontent

import "fmt"

type Context struct {
	FileContent *FileContent
	Line        int      // 1-
	Index       int      // 0- FileBox.Contet の Index
	Offset      int      // 0 マクロ未展開, 1- マクロ展開後
	Source      *Context // マクロ定義 Context
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
	return c.FileContent == o.FileContent && c.Line == o.Line && c.Index == o.Index && c.Offset == o.Offset
}

// Context の示す行を取得
func (c *Context) GetLine() string {
	if c == nil {
		return "<Context nil>"
	}
	if c.Offset == 0 {
		s, _ := c.FileContent.GetLine(c.Line)
		return s
	}
	if c.Source != nil {
		s, _ := c.Source.FileContent.GetLine(c.Source.Line)
		return s
	}
	return "<Context.Source nil>"
}
