package web

import "github.com/life4/gweb/canvas"

type Canvas struct {
	HTMLElement
}

// getters

func (element Canvas) Context(name string) canvas.Context {
	value := element.Call("getContext", name)
	return canvas.Context{Value: value.JSValue()}
}

func (element Canvas) Context2D() canvas.Context2D {
	context := element.Context("2d")
	return context.Context2D()
}

func (element Canvas) Width() int {
	return element.Get("width").Int()
}

func (element Canvas) Height() int {
	return element.Get("height").Int()
}

// setters

func (element Canvas) SetWidth(value int) {
	element.Set("width", value)
}

func (element Canvas) SetHeight(value int) {
	element.Set("height", value)
}
