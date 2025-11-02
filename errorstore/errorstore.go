package errorstore

import "fmt"

type ErrorMessage struct {
	Message    string
	Filename   string
	LineNumber int
}

func (e ErrorMessage) String() string {
	return fmt.Sprintf("%s:%d:%s", e.Filename, e.LineNumber, e.Message)
}

type ErrorStore struct {
	Errors   []ErrorMessage
	Warnings []ErrorMessage
}

func New() *ErrorStore {
	return &ErrorStore{}
}

func (es *ErrorStore) AddError(msg string, file string, line int) {
	es.Errors = append(es.Errors, ErrorMessage{file, msg, line})
}

func (es *ErrorStore) AddWarning(msg string, file string, line int) {
	es.Errors = append(es.Errors, ErrorMessage{msg, file, line})
}

func (es *ErrorStore) Count() (int, int) {
	return len(es.Errors), len(es.Warnings)
}

func (es *ErrorStore) Print() {
	fmt.Printf("%d errros\n", len(es.Errors))
	for _, e := range es.Errors {
		fmt.Println(e.String())
	}
	fmt.Printf("%d warnings\n", len(es.Warnings))
	for _, e := range es.Warnings {
		fmt.Println(e.String())
	}
}
