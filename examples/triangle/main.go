package main

import "github.com/life4/gweb/web"

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Canvas triangle example")
	body := doc.Body()

	// create canvas
	h := window.InnerHeight() - 40
	w := window.InnerWidth() - 40
	canvas := doc.CreateCanvas()
	canvas.SetHeight(h)
	canvas.SetWidth(w)
	body.Node().AppendChild(canvas.Node())

	context := canvas.Context2D()

	// draw black background
	context.SetFillStyle("black")
	context.BeginPath()
	context.Rect(0, 0, w, h)
	context.Fill()
	context.ClosePath()

	// draw red triangle
	centerX := w / 2
	centerY := h / 2
	context.SetFillStyle("red")
	context.BeginPath()
	context.MoveTo(centerX-40, centerY+40)
	context.LineTo(centerX+40, centerY+40)
	context.LineTo(centerX, centerY-40)
	context.Fill()
	context.ClosePath()
}
