package fileblock

import "fmt"

type Context struct {
	FileBlock *FileBlock
	Line      int // 1-
	Column    int // 0-
}

func (c *Context) String() string {
	return fmt.Sprintf("%q:(%d,%d)", c.FileBlock.Filename, c.Line, c.Column)
}
