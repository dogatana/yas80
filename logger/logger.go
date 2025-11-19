package logger

import "fmt"

type ErrorMessage struct {
	Message    string
	Filename   string
	LineNumber int
}

func (em *ErrorMessage) Error() string {
	return fmt.Sprintf("%q:%d [ERROR] %s", em.Filename, em.LineNumber, em.Message)
}

type WarningMessage struct {
	Message    string
	Filename   string
	LineNumber int
}

func (wm WarningMessage) Error() string {
	return fmt.Sprintf("%q:%d [WARN] %s", wm.Filename, wm.LineNumber, wm.Message)
}

type InfoMessage struct {
	Message    string
	Filename   string
	LineNumber int
}

func (im InfoMessage) Error() string {
	return fmt.Sprintf("%q:%d [INFO] %s", im.Filename, im.LineNumber, im.Message)
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

func (l *Logger) Error(msg string, line int) error {
	err := &ErrorMessage{msg, l.Filename, line}
	l.Errors = append(l.Errors, err)
	return err
}

func (l *Logger) Warning(msg string, line int) error {
	err := &WarningMessage{msg, l.Filename, line}
	l.Warnings = append(l.Warnings, err)
	return err
}
func (l *Logger) Info(msg string, line int) error {
	err := &InfoMessage{msg, l.Filename, line}
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
