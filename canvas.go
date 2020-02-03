package glowasm

import (
	"syscall/js"
)

type Canvas struct {
	value js.Value
}

// getters

func (canvas *Canvas) GetContext(name string) Value {
	value := canvas.value.Call("getContext")
	return Value{value: value}
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
