package log

import (
	"bytes"
	"fmt"
	"log"
	"sort"
)

// TODO(casey): Should handle the default case, we now need a GetLevel function
// to be considered a handler.

type defaultHandler struct {
	Level: Level
}

// field used for sorting.
type field struct {
	Name  string
	Value interface{}
}

// by sorts fields by name.
type byName []field

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func NewDefaultHandler() *defaultHandler {
	return &defaultHandler{
		Level: InfoLevel
	}
}

func (h *defaultHandler) GetLevel() Level {
	return InfoLevel
}

// handleStdLog outputs to the stlib log.
func (h *defaultHandler) HandleLog(e *Entry) error {
	level := levelNames[e.Level]

	var fields []field

	for k, v := range e.Fields {
		fields = append(fields, field{k, v})
	}

	sort.Sort(byName(fields))

	var b bytes.Buffer
	fmt.Fprintf(&b, "%5s %-25s", level, e.Message)

	for _, f := range fields {
		fmt.Fprintf(&b, " %s=%v", f.Name, f.Value)
	}

	log.Println(b.String())

	return nil
}
