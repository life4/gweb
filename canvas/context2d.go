package canvas

import "syscall/js"

type Context2D struct {
	value js.Value
}

// SUBTYPE GETTERS

func (context Context2D) Shadow() Shadow {
	return Shadow(context)
}

func (context Context2D) Line() Line {
	return Line(context)
}

func (context Context2D) Rectangle(x, y, width, height int) Rectangle {
	return Rectangle{value: context.value, x: x, y: y, width: width, height: height}
}

func (context Context2D) Text() Text {
	return Text(context)
}

// STYLES

func (context Context2D) FillStyle() string {
	return context.value.Get("fillStyle").String()
}

func (context Context2D) SetFillStyle(value string) {
	context.value.Set("fillStyle", value)
}

func (context Context2D) StrokeStyle() string {
	return context.value.Get("strokeStyle").String()
}

func (context Context2D) SetStrokeStyle(value string) {
	context.value.Set("strokeStyle", value)
}

// OTHER ATTRS

// GlobalAlpha returns the current alpha or transparency value of the drawing
func (context Context2D) GlobalAlpha() float64 {
	return context.value.Get("globalAlpha").Float()
}

func (context Context2D) SetGlobalAlpha(value float64) {
	context.value.Set("globalAlpha", value)
}

func (context Context2D) GlobalCompositeOperation() string {
	return context.value.Get("globalCompositeOperation").String()
}

func (context Context2D) SetGlobalCompositeOperation(value string) {
	context.value.Set("globalCompositeOperation", value)
}

// PATH API

func (context Context2D) BeginPath() {
	context.value.Call("beginPath")
}

func (context Context2D) ClosePath() {
	context.value.Call("closePath")
}

func (context Context2D) Arc(x, y, r int, sAngle, eAngle float64) {
	context.value.Call("arc", x, y, r, sAngle, eAngle, false)
}

func (context Context2D) ArcTo(x1, y1, x2, y2, r int) {
	context.value.Call("arcTo", x1, y1, x2, y2, r)
}

func (context Context2D) Clip() {
	context.value.Call("clip")
}

func (context Context2D) Fill() {
	context.value.Call("fill")
}

// IsPointInPath returns true if the specified point is in the current path
func (context Context2D) IsPointInPath(x int, y int) bool {
	return context.value.Call("isPointInPath", x, y).Bool()
}

func (context Context2D) LineTo(x int, y int) {
	context.value.Call("lineTo", x, y)
}

func (context Context2D) MoveTo(x int, y int) {
	context.value.Call("moveTo", x, y)
}

func (context Context2D) Stroke() {
	context.value.Call("stroke")
}

// BezierCurveTo adds a point to the current path by using the specified
// control points that represent a cubic Bézier curve.
func (context Context2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y int) {
	context.value.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

// QuadraticCurveTo adds a point to the current path by using the specified
// control points that represent a quadratic Bézier curve.
func (context Context2D) QuadraticCurveTo(cpx, cpy, x, y int) {
	context.value.Call("quadraticCurveTo", cpx, cpy, x, y)
}

// TRANSFORMATION API

// Rotate rotates the current drawing
func (context Context2D) Rotate(angle float64) {
	context.value.Call("scale", angle)
}

// Scale scales the current drawing bigger or smaller
func (context Context2D) Scale(x float64, y float64) {
	context.value.Call("scale", x, y)
}

// Transform replaces the current transformation matrix.
// a: Horizontal scaling
// b: Horizontal skewing
// c: Vertical skewing
// d: Vertical scaling
// e: Horizontal moving
// f: Vertical moving
func (context Context2D) Transform(a, b, c, d, e, f float64) {
	context.value.Call("transform", a, b, c, d, e, f)
}

// Translate remaps the (0,0) position on the canvas
func (context Context2D) Translate(x float64, y float64) {
	context.value.Call("translate", x, y)
}
