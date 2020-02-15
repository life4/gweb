package canvas

import "syscall/js"

type Shadow struct {
	value js.Value
}

func (context Shadow) Blur() float64 {
	return context.value.Get("shadowBlur").Float()
}

func (context Shadow) SetBlur(value float64) {
	context.value.Set("shadowBlur", value)
}

func (context Shadow) Color() string {
	return context.value.Get("shadowColor").String()
}

func (context Shadow) SetColor(value string) {
	context.value.Set("shadowColor", value)
}

func (context Shadow) OffsetX() float64 {
	return context.value.Get("shadowOffsetX").Float()
}

func (context Shadow) SetOffsetX(value float64) {
	context.value.Set("shadowOffsetX", value)
}

func (context Shadow) OffsetY() float64 {
	return context.value.Get("shadowOffsetY").Float()
}

func (context Shadow) SetOffsetY(value float64) {
	context.value.Set("shadowOffsetY", value)
}
