package canvas

import "syscall/js"

type Line struct {
	value js.Value
}

func (context Line) Cap() string {
	return context.value.Get("lineCap").String()
}

func (context Line) SetCap(value string) {
	context.value.Set("lineCap", value)
}

func (context Line) Join() string {
	return context.value.Get("lineJoin").String()
}

func (context Line) SetJoin(value string) {
	context.value.Set("lineJoin", value)
}

func (context Line) MiterLimit() string {
	return context.value.Get("miterLimit").String()
}

func (context Line) SetMiterLimit(value string) {
	context.value.Set("miterLimit", value)
}

func (context Line) Width() int {
	return context.value.Get("lineWidth").Int()
}

func (context Line) SetWidth(value int) {
	context.value.Set("lineWidth", value)
}

func (context Line) Draw(x1, y1, x2, y2 int) {
	context.value.Call("beginPath")
	context.value.Call("moveTo", x1, y1)
	context.value.Call("lineTo", x2, y2)
	context.value.Call("stroke")
}
