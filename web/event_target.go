package web

import (
	"syscall/js"
)

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
type EventTarget struct {
	Value
}

// Listen registers callback for the given event.
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (target EventTarget) Listen(event EventType, handler func(event Event)) {
	wrapped := func(this js.Value, args []js.Value) any {
		v := Value{Value: args[0]}
		handler(v.Event())
		return nil
	}
	target.Call("addEventListener", string(event), js.FuncOf(wrapped))
}
