package web

import (
	"syscall/js"
)

type EventTarget struct {
	Value
}

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (target EventTarget) Listen(event EventType, handler func(event Event)) {
	wrapped := func(this js.Value, args []js.Value) interface{} {
		v := Value{Value: args[0]}
		handler(v.Event())
		return nil
	}
	target.Call("addEventListener", event, js.FuncOf(wrapped))
}
