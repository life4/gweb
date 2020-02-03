package glowasm

import (
	"syscall/js"
)

type Canvas struct {
	value js.Value
}

// getters

func (canvas *Canvas) Context(name string) RenderingContext {
	value := canvas.value.Call("getContext")
	return RenderingContext{value: value}
}

func (canvas *Canvas) Context2D(name string) Context2D {
	context := canvas.Context("2d")
	return context.Context2D()
}

func (canvas *Canvas) Width() int {
	return canvas.value.Get("width").Int()
}

func (canvas *Canvas) Height() int {
	return canvas.value.Get("height").Int()
}

// setters

func (canvas *Canvas) SetWidth(value int) {
	canvas.value.Set("width", value)
}

func (canvas *Canvas) SetHeight(value int) {
	canvas.value.Set("height", value)
}
