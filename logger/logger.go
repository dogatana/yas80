package logger

import "fmt"

type ErrorMessage struct {
	Message    string
	Filename   string
	LineNumber int
}

func (em ErrorMessage) String() string {
	return fmt.Sprintf("%q:%d [ERROR] %s", em.Filename, em.LineNumber, em.Message)
}

type WarningMessage struct {
	Message    string
	Filename   string
	LineNumber int
}

func (wm WarningMessage) String() string {
	return fmt.Sprintf("%q:%d [WARN] %s", wm.Filename, wm.LineNumber, wm.Message)
}

type InfoMessage struct {
	Message    string
	Filename   string
	LineNumber int
}

func (im InfoMessage) String() string {
	return fmt.Sprintf("%q:%d [WARN] %s", im.Filename, im.LineNumber, im.Message)
}

type Logger struct {
	Errors     []ErrorMessage
	Warnings   []WarningMessage
	Infomation []InfoMessage
	Filename   string
}

func New(filename string) *Logger {
	return &Logger{Filename: filename}
}

func (l *Logger) Error(msg string, line int) {
	l.Errors = append(l.Errors, ErrorMessage{msg, l.Filename, line})
}

func (l *Logger) Warning(msg string, line int) {
	l.Warnings = append(l.Warnings, WarningMessage{msg, l.Filename, line})
}
func (l *Logger) Info(msg string, line int) {
	l.Infomation = append(l.Infomation, InfoMessage{msg, l.Filename, line})
}

func (l *Logger) Count() (int, int) {
	return len(l.Errors), len(l.Warnings)
}

func (l *Logger) Print() {
	fmt.Printf("%d errros\n", len(l.Errors))
	for _, e := range l.Errors {
		fmt.Println(e.String())
	}
	fmt.Printf("%d warnings\n", len(l.Warnings))
	for _, e := range l.Warnings {
		fmt.Println(e.String())
	}
	fmt.Printf("%d info\n", len(l.Infomation))
	for _, e := range l.Infomation {
		fmt.Println(e.String())
	}
}
