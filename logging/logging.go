package logging

import (
	"fmt"
	"yas80/filecontent"
	"yas80/internal/util"
)

type MessageType int

const (
	Err = iota
	Warn
	Info
)

var msgTypeName = map[MessageType]string{
	Err:  "[ERR]",
	Warn: "[WARN]",
	Info: "[INFO]",
}

type Message struct {
	Type    MessageType
	Text    string
	Context *filecontent.Context
}

func (m *Message) Message() string { return msgTypeName[m.Type] + m.Text }
func (m *Message) String() string {
	tn := msgTypeName[m.Type]
	if m.Context == nil {
		return fmt.Sprintf("%s %s", tn, m.Text)
	} else {
		return fmt.Sprintf("%q:%d %s %s", m.Context.FileContent.Filename, m.Context.Line, tn, m.Text)
	}
}

func (m *Message) LString() string { // Lister 用
	tn := msgTypeName[m.Type]
	return tn + " " + m.Text
}

// 2 つのメッセージが等しいかどうか（メソッド）
func (m *Message) Equal(o *Message) bool {
	return Equal(m, o)
}

// 2 つのメッセージが等しいかどうか（関数）
func Equal(a, b *Message) bool {
	return a.Type == b.Type && a.Text == b.Text && a.Context.Equal(b.Context)
}

type Logger struct {
	messages []*Message
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Count() (int, int, int) {
	var e, w, i int
	for _, m := range l.messages {
		switch m.Type {
		case Err:
			e++
		case Warn:
			w++
		case Info:
			i++
		}
	}
	return e, w, i
}

func (l *Logger) ErrorCount() int {
	e, _, _ := l.Count()
	return e
}

// 全メッセージ取得
func (l *Logger) GetMessages() []*Message {
	return l.messages
}

// Error メッセージ取得
func (l *Logger) GetErrors() []*Message {
	return util.Filter(l.messages, func(m *Message) bool { return m.Type == Err })
}

// Warning メッセージ取得
func (l *Logger) GetWarnings() []*Message {
	return util.Filter(l.messages, func(m *Message) bool { return m.Type == Warn })
}

// Information メッセージ取得
func (l *Logger) GetInformation() []*Message {
	return util.Filter(l.messages, func(m *Message) bool { return m.Type == Info })
}

// Error 追加
func (l *Logger) Error(msg string, ctx *filecontent.Context) *Message {
	m := &Message{Type: Err, Text: msg, Context: ctx}
	l.messages = append(l.messages, m)
	return m
}

// Warning 追加
func (l *Logger) Warning(msg string, ctx *filecontent.Context) *Message {
	m := &Message{Type: Warn, Text: msg, Context: ctx}
	l.messages = append(l.messages, m)
	return m
}

// Information 追加
func (l *Logger) Info(msg string, ctx *filecontent.Context) *Message {
	m := &Message{Type: Info, Text: msg, Context: ctx}
	l.messages = append(l.messages, m)
	return m
}

// logMessage の表示
func (l *Logger) Print() {
	ec, wc, ic := l.Count()
	// Warning, Information の重複を削除
	if ec != 0 {
		fmt.Printf("%d errros\n", ec)
		for _, m := range util.Filter(l.messages, func(m *Message) bool { return m.Type == Err }) {
			fmt.Println(m)
		}
	}
	if wc != 0 {
		fmt.Printf("%d warnings\n", wc)
		for _, m := range util.Filter(l.messages, func(m *Message) bool { return m.Type == Warn }) {
			fmt.Println(m)
		}
	}
	if ic != 0 {
		fmt.Printf("%d information\n", ic)
		for _, m := range util.Filter(l.messages, func(m *Message) bool { return m.Type == Info }) {
			fmt.Println(m)
		}
	}
}

// 重複メッセージの削除
func (l *Logger) RemoveDupe() {
	if len(l.messages) < 2 {
		return
	}
	for i := 0; i < len(l.messages)-1; i++ {
		m := l.messages[i]
		if m == nil {
			continue
		}
		for j := i + 1; j < len(l.messages); j++ {
			if l.messages[j] == nil {
				continue
			}
			if m.Equal(l.messages[j]) {
				l.messages[j] = nil
			}
		}
	}
	l.messages = util.Filter(l.messages, func(m *Message) bool { return m != nil })
}
