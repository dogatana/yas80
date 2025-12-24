package fileblock

import "fmt"

type Context struct {
	FileBlock *FileBlock
	Line      int // 1-
	Index     int // 0- FileBox.Contet の Index
}

func (c *Context) String() string {
	return fmt.Sprintf("%q:(%d,%d)", c.FileBlock.Filename, c.Line, c.Index)
}

func (c *Context) Equal(other *Context) bool {
	return c.FileBlock == other.FileBlock && c.Line == other.Line && c.Index == other.Index
}
