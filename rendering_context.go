package glowasm

import (
	"syscall/js"
)

type RenderingContext struct {
	value js.Value
}

func (context *RenderingContext) Context2D() Context2D {
	return Context2D{value: context.value}
}

type Context2D struct {
	value js.Value
}

// STYLES

func (context *Context2D) FillStyle() string {
	return context.value.Get("fillStyle").String()
}

func (context *Context2D) SetFillStyle(value string) {
	context.value.Set("fillStyle", value)
}

func (context *Context2D) StrokeStyle() string {
	return context.value.Get("strokeStyle").String()
}

func (context *Context2D) SetStrokeStyle(value string) {
	context.value.Set("strokeStyle", value)
}

func (context *Context2D) Shadow() Shadow {
	return Shadow{value: context.value}
}

// PATH API

func (context *Context2D) BeginPath() {
	context.value.Call("beginPath")
}

func (context *Context2D) ClosePath() {
	context.value.Call("closePath")
}

func (context *Context2D) Arc(x, y, r int, sAngle, eAngle float64) {
	context.value.Call("arc", x, y, r, sAngle, eAngle)
}

func (context *Context2D) ArcTo(x1, y1, x2, y2, r int) {
	context.value.Call("arcTo", x1, y1, x2, y2, r)
}

func (context *Context2D) Clip() {
	context.value.Call("clip")
}

func (context *Context2D) Fill() {
	context.value.Call("fill")
}

// IsPointInPath returns true if the specified point is in the current path
func (context *Context2D) IsPointInPath(x int, y int) bool {
	return context.value.Call("isPointInPath", x, y).Bool()
}

func (context *Context2D) LineTo(x int, y int) {
	context.value.Call("lineTo", x, y)
}

func (context *Context2D) MoveTo(x int, y int) {
	context.value.Call("moveTo", x, y)
}

func (context *Context2D) Stroke() {
	context.value.Call("stroke")
}

// BezierCurveTo adds a point to the current path by using the specified
// control points that represent a cubic Bézier curve.
func (context *Context2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y int) {
	context.value.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

// QuadraticCurveTo adds a point to the current path by using the specified
// control points that represent a quadratic Bézier curve.
func (context *Context2D) QuadraticCurveTo(cpx, cpy, x, y int) {
	context.value.Call("quadraticCurveTo", cpx, cpy, x, y)
}

// RECTANGLES

func (context *Context2D) Rect(x, y, width, height int) {
	context.value.Call("rect", x, y, width, height)
}

func (context *Context2D) FillRect(x, y, width, height int) {
	context.value.Call("fillRect", x, y, width, height)
}

func (context *Context2D) StrokeRect(x, y, width, height int) {
	context.value.Call("strokeRect", x, y, width, height)
}

func (context *Context2D) ClearRect(x, y, width, height int) {
	context.value.Call("clearRect", x, y, width, height)
}

// TRANSFORMATION API

// Rotate rotates the current drawing
func (context *Context2D) Rotate(angle float64) {
	context.value.Call("scale", angle)
}

// Scale scales the current drawing bigger or smaller
func (context *Context2D) Scale(x float64, y float64) {
	context.value.Call("scale", x, y)
}

// Transform replaces the current transformation matrix.
// a: Horizontal scaling
// b: Horizontal skewing
// c: Vertical skewing
// d: Vertical scaling
// e: Horizontal moving
// f: Vertical moving
func (context *Context2D) Transform(a, b, c, d, e, f float64) {
	context.value.Call("transform", a, b, c, d, e, f)
}

// Translate remaps the (0,0) position on the canvas
func (context *Context2D) Translate(x float64, y float64) {
	context.value.Call("translate", x, y)
}

// CONTEXT SUBTYPES

type Shadow struct {
	value js.Value
}

func (shadow *Shadow) Blur() float64 {
	return shadow.value.Get("shadowBlur").Float()
}

func (shadow *Shadow) Color() string {
	return shadow.value.Get("shadowColor").String()
}

func (shadow *Shadow) OffsetX() float64 {
	return shadow.value.Get("shadowOffsetX").Float()
}

func (shadow *Shadow) OffsetY() float64 {
	return shadow.value.Get("shadowOffsetY").Float()
}

type Line struct {
	value js.Value
}

func (line *Line) Cap() string {
	return line.value.Get("lineCap").String()
}

func (line *Line) Join() string {
	return line.value.Get("lineJoin").String()
}

func (line *Line) MiterLimit() string {
	return line.value.Get("miterLimit").String()
}

func (line *Line) Width() int {
	return line.value.Get("lineWidth").Int()
}
