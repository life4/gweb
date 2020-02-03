package glowasm

import (
	"syscall/js"
)

type Canvas struct {
	value js.Value
}

func (canvas *Canvas) GetContext(name string) Value {
	value := canvas.value.Call("getContext")
	return Value{value: value}
}

// CONTEXT

type RenderingContext struct {
	value js.Value
}

func (context *RenderingContext) RenderingContext2D() RenderingContext2D {
	return RenderingContext2D{value: context.value}
}

type RenderingContext2D struct {
	value js.Value
}
