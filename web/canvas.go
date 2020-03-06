package web

import "github.com/life4/gweb/canvas"

// Canvas provides properties and methods for manipulating the layout and presentation of <canvas> elements.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement
type Canvas struct {
	HTMLElement
}

// getters

// Context returns a drawing context on the canvas, or null if the context ID is not supported.
// A drawing context lets you draw on the canvas.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/getContext
func (element Canvas) Context(name string) canvas.Context {
	value := element.Call("getContext", name)
	return canvas.Context{Value: value.JSValue()}
}

// Context2D returns 2D context to draw on canvas.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/getContext
func (element Canvas) Context2D() canvas.Context2D {
	context := element.Context("2d")
	return context.Context2D()
}

// Width is the width of the <canvas> element interpreted in CSS pixels
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/width
func (element Canvas) Width() int {
	return element.Get("width").Int()
}

// Height is the height of the <canvas> element interpreted in CSS pixels
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/height
func (element Canvas) Height() int {
	return element.Get("height").Int()
}

// setters

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/width
func (element Canvas) SetWidth(value int) {
	element.Set("width", value)
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/height
func (element Canvas) SetHeight(value int) {
	element.Set("height", value)
}
