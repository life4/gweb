package glowasm

type Canvas struct {
	Element
}

// getters

func (canvas *Canvas) Context(name string) RenderingContext {
	value := canvas.Call("getContext")
	return RenderingContext{Value: value}
}

func (canvas *Canvas) Context2D(name string) Context2D {
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
