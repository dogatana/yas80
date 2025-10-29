package errorstore

import "fmt"

type ErrorMessage struct {
	Filename   string
	LineNumber int
	Message    string
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

func (es *ErrorStore) AddError(file string, line int, msg string) {
	es.Errors = append(es.Errors, ErrorMessage{file, line, msg})
}

func (es *ErrorStore) AddWarning(file string, line int, msg string) {
	es.Errors = append(es.Errors, ErrorMessage{file, line, msg})
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
