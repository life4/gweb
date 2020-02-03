package glowasm

import (
	"syscall/js"
)

type RenderingContext struct {
	value js.Value
}

func (context *RenderingContext) RenderingContext2D() RenderingContext2D {
	return RenderingContext2D{value: context.value}
}

type RenderingContext2D struct {
	value js.Value
}

// STYLES

func (context *RenderingContext2D) FillStyle() string {
	return context.value.Get("fillStyle").String()
}

func (context *RenderingContext2D) SetFillStyle(value string) {
	context.value.Set("fillStyle", value)
}

func (context *RenderingContext2D) StrokeStyle() string {
	return context.value.Get("strokeStyle").String()
}

func (context *RenderingContext2D) SetStrokeStyle(value string) {
	context.value.Set("strokeStyle", value)
}

func (context *RenderingContext2D) Shadow() Shadow {
	return Shadow{value: context.value}
}

// PATH API

func (context *RenderingContext2D) BeginPath() {
	context.value.Call("beginPath")
}

func (context *RenderingContext2D) ClosePath() {
	context.value.Call("closePath")
}

func (context *RenderingContext2D) Arc(x, y, r int, sAngle, eAngle float64) {
	context.value.Call("arc", x, y, r, sAngle, eAngle)
}

func (context *RenderingContext2D) ArcTo(x1, y1, x2, y2, r int) {
	context.value.Call("arcTo", x1, y1, x2, y2, r)
}

func (context *RenderingContext2D) Clip() {
	context.value.Call("clip")
}

func (context *RenderingContext2D) Fill() {
	context.value.Call("fill")
}

// IsPointInPath returns true if the specified point is in the current path
func (context *RenderingContext2D) IsPointInPath(x int, y int) bool {
	return context.value.Call("isPointInPath", x, y).Bool()
}

func (context *RenderingContext2D) LineTo(x int, y int) {
	context.value.Call("lineTo", x, y)
}

func (context *RenderingContext2D) MoveTo(x int, y int) {
	context.value.Call("moveTo", x, y)
}

func (context *RenderingContext2D) Stroke() {
	context.value.Call("stroke")
}

// BezierCurveTo adds a point to the current path by using the specified
// control points that represent a cubic Bézier curve.
func (context *RenderingContext2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y int) {
	context.value.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

// QuadraticCurveTo adds a point to the current path by using the specified
// control points that represent a quadratic Bézier curve.
func (context *RenderingContext2D) QuadraticCurveTo(cpx, cpy, x, y int) {
	context.value.Call("quadraticCurveTo", cpx, cpy, x, y)
}

// TRANSFORMATION API

// Rotate rotates the current drawing
func (context *RenderingContext2D) Rotate(angle float64) {
	context.value.Call("scale", angle)
}

// Scale scales the current drawing bigger or smaller
func (context *RenderingContext2D) Scale(x float64, y float64) {
	context.value.Call("scale", x, y)
}

// Transform replaces the current transformation matrix.
// a: Horizontal scaling
// b: Horizontal skewing
// c: Vertical skewing
// d: Vertical scaling
// e: Horizontal moving
// f: Vertical moving
func (context *RenderingContext2D) Transform(a, b, c, d, e, f float64) {
	context.value.Call("transform", a, b, c, d, e, f)
}

// Translate remaps the (0,0) position on the canvas
func (context *RenderingContext2D) Translate(x float64, y float64) {
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
