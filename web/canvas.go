package web

type Canvas struct {
	HTMLElement
}

// getters

func (canvas *Canvas) Context(name string) RenderingContext {
	value := canvas.Call("getContext", name)
	return RenderingContext{Value: value}
}

func (canvas *Canvas) Context2D() Context2D {
	context := canvas.Context("2d")
	return context.Context2D()
}

func (canvas *Canvas) Width() int {
	return canvas.Get("width").Int()
}

func (canvas *Canvas) Height() int {
	return canvas.Get("height").Int()
}

// setters

func (canvas *Canvas) SetWidth(value int) {
	canvas.Set("width", value)
}

func (canvas *Canvas) SetHeight(value int) {
	canvas.Set("height", value)
}
