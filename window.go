package glowasm

import (
	"syscall/js"
)

type Window struct {
	Value
}

func GetWindow() Window {
	value := Value{Value: js.Global()}
	return Window{Value: value}
}

// SUBTYPE GETTERS

func (doc *Document) Document() Document {
	return Document{Value: doc.Value}
}
