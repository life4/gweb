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

func (window Window) Console() Console {
	return Console{Value: window.Get("console")}
}

func (window Window) Document() Document {
	return Document{Value: window.Get("document")}
}
