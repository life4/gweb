package web

type RenderingContext struct {
	Value
}

func (context *RenderingContext) Context2D() Context2D {
	return Context2D{value: context.Value}
}

type Context2D struct {
	value Value
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

// OTHER ATTRS

// GlobalAlpha returns the current alpha or transparency value of the drawing
func (context *Context2D) GlobalAlpha() float64 {
	return context.value.Get("globalAlpha").Float()
}

func (context *Context2D) SetGlobalAlpha(value float64) {
	context.value.Set("globalAlpha", value)
}

func (context *Context2D) GlobalCompositeOperation() string {
	return context.value.Get("globalCompositeOperation").String()
}

func (context *Context2D) SetGlobalCompositeOperation(value string) {
	context.value.Set("globalCompositeOperation", value)
}

// PATH API

func (context *Context2D) BeginPath() {
	context.value.Call("beginPath")
}

func (context *Context2D) ClosePath() {
	context.value.Call("closePath")
}

func (context *Context2D) Arc(x, y, r int, sAngle, eAngle float64) {
	context.value.Call("arc", x, y, r, sAngle, eAngle, false)
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

// HELPERS

func (context Context2D) RoundedRect(x, y, width, height, radius int) {
	context.BeginPath()
	context.MoveTo(x, y+radius)
	context.LineTo(x, y+height-radius)
	context.ArcTo(x, y+height, x+radius, y+height, radius)
	context.LineTo(x+width-radius, y+height)
	context.ArcTo(x+width, y+height, x+width, y+height-radius, radius)
	context.LineTo(x+width, y+radius)
	context.ArcTo(x+width, y, x+width-radius, y, radius)
	context.LineTo(x+radius, y)
	context.ArcTo(x, y, x, y+radius, radius)
	context.Stroke()
}

// CONTEXT SUBTYPES

type Shadow struct {
	value Value
}

func (context *Shadow) Blur() float64 {
	return context.value.Get("shadowBlur").Float()
}

func (context *Shadow) SetBlur(value float64) {
	context.value.Set("shadowBlur", value)
}

func (context *Shadow) Color() string {
	return context.value.Get("shadowColor").String()
}

func (context *Shadow) SetColor(value string) {
	context.value.Set("shadowColor", value)
}

func (context *Shadow) OffsetX() float64 {
	return context.value.Get("shadowOffsetX").Float()
}

func (context *Shadow) SetOffsetX(value float64) {
	context.value.Set("shadowOffsetX", value)
}

func (context *Shadow) OffsetY() float64 {
	return context.value.Get("shadowOffsetY").Float()
}

func (context *Shadow) SetOffsetY(value float64) {
	context.value.Set("shadowOffsetY", value)
}

type Line struct {
	value Value
}

func (context *Line) Cap() string {
	return context.value.Get("lineCap").String()
}

func (context *Line) SetCap(value string) {
	context.value.Set("lineCap", value)
}

func (context *Line) Join() string {
	return context.value.Get("lineJoin").String()
}

func (context *Line) SetJoin(value string) {
	context.value.Set("lineJoin", value)
}

func (context *Line) MiterLimit() string {
	return context.value.Get("miterLimit").String()
}

func (context *Line) SetMiterLimit(value string) {
	context.value.Set("miterLimit", value)
}

func (context *Line) Width() int {
	return context.value.Get("lineWidth").Int()
}

func (context *Line) SetWidth(value int) {
	context.value.Set("lineWidth", value)
}

func (context *Line) Draw(x1, y1, x2, y2 int) {
	context.value.Call("beginPath")
	context.value.Call("moveTo", x1, y1)
	context.value.Call("lineTo", x2, y2)
	context.value.Call("stroke")
}

type Text struct {
	value Value
}

func (context *Text) Align() string {
	return context.value.Get("align").String()
}

func (context *Line) SetAlign(value string) {
	context.value.Set("align", value)
}

func (context *Text) Baseline() string {
	return context.value.Get("baseline").String()
}

func (context *Line) SetBaseline(value string) {
	context.value.Set("baseline", value)
}

func (context *Text) Font() string {
	return context.value.Get("font").String()
}

func (context *Line) SetFont(value string) {
	context.value.Set("font", value)
}

func (context *Text) Fill(text string, x, y, maxWidth int) {
	if maxWidth == 0 {
		context.value.Call("fillText", x, y)
	} else {
		context.value.Call("fillText", x, y, maxWidth)
	}
}

func (context *Text) Stroke(text string, x, y, maxWidth int) {
	if maxWidth == 0 {
		context.value.Call("strokeText", x, y)
	} else {
		context.value.Call("strokeText", x, y, maxWidth)
	}
}

func (context *Text) Width(text string) int {
	return context.value.Call("measureText", text).Get("width").Int()
}
