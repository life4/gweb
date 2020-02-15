package canvas

import "syscall/js"

type Text struct {
	value js.Value
}

func (context Text) Align() string {
	return context.value.Get("align").String()
}

func (context Text) SetAlign(value string) {
	context.value.Set("align", value)
}

func (context Text) Baseline() string {
	return context.value.Get("baseline").String()
}

func (context Text) SetBaseline(value string) {
	context.value.Set("baseline", value)
}

func (context Text) Font() string {
	return context.value.Get("font").String()
}

func (context Text) SetFont(value string) {
	context.value.Set("font", value)
}

func (context Text) Fill(text string, x, y, maxWidth int) {
	if maxWidth <= 0 {
		context.value.Call("fillText", text, x, y)
	} else {
		context.value.Call("fillText", text, x, y, maxWidth)
	}
}

func (context Text) Stroke(text string, x, y, maxWidth int) {
	if maxWidth == 0 {
		context.value.Call("strokeText", x, y)
	} else {
		context.value.Call("strokeText", x, y, maxWidth)
	}
}

func (context Text) Width(text string) int {
	return context.value.Call("measureText", text).Get("width").Int()
}
