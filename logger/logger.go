package logger

import (
	"fmt"
	"yas80/fileblock"
)

type ErrorMessage struct {
	Message string
	Context *fileblock.Context
}

func (em *ErrorMessage) Error() string {
	return fmt.Sprintf("%q:%d [ERROR] %s", em.Context.FileBlock.Filename, em.Context.Line, em.Message)
}

type WarningMessage struct {
	Message string
	Context *fileblock.Context
}

func (wm WarningMessage) Error() string {
	return fmt.Sprintf("%q:%d [WARN] %s", wm.Context.FileBlock.Filename, wm.Context.Line, wm.Message)
}

type InfoMessage struct {
	Message string
	Context *fileblock.Context
}

func (im InfoMessage) Error() string {
	return fmt.Sprintf("%q:%d [INFO] %s", im.Context.FileBlock.Filename, im.Context.Line, im.Message)
}

type Logger struct {
	Errors     []*ErrorMessage
	Warnings   []*WarningMessage
	Infomation []*InfoMessage
	Filename   string
}

func New(filename string) *Logger {
	return &Logger{Filename: filename}
}

func (l *Logger) Error(msg string, ctx *fileblock.Context) error {
	err := &ErrorMessage{Message: msg, Context: ctx}
	l.Errors = append(l.Errors, err)
	return err
}

func (l *Logger) Warning(msg string, ctx *fileblock.Context) error {
	err := &WarningMessage{Message: msg, Context: ctx}
	l.Warnings = append(l.Warnings, err)
	return err
}
func (l *Logger) Info(msg string, ctx *fileblock.Context) error {
	err := &InfoMessage{Message: msg, Context: ctx}
	l.Infomation = append(l.Infomation, err)
	return err
}

func (l *Logger) Count() (int, int, int) {
	return len(l.Errors), len(l.Warnings), len(l.Infomation)
}

func (l *Logger) Print() {
	if len(l.Errors) != 0 {
		fmt.Printf("%d errros\n", len(l.Errors))
		for _, e := range l.Errors {
			fmt.Println(e.Error())
		}
	}
	if len(l.Warnings) != 0 {
		fmt.Printf("%d warnings\n", len(l.Warnings))
		for _, e := range l.Warnings {
			fmt.Println(e.Error())
		}
	}
	if len(l.Infomation) != 0 {
		fmt.Printf("%d info\n", len(l.Infomation))
		for _, e := range l.Infomation {
			fmt.Println(e.Error())
		}
	}
}
