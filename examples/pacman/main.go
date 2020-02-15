package main

import (
	"math"

	"github.com/life4/gweb/web"
)

// https://developer.mozilla.org/en-US/docs/Web/API/Canvas_API/Tutorial/Drawing_shapes#Making_combinations
func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Canvas pacman example")
	body := doc.Body()

	// create canvas
	canvas := doc.CreateCanvas()
	canvas.SetHeight(150)
	canvas.SetWidth(150)
	body.Node().AppendChild(canvas.Node())

	ctx := canvas.Context2D()

	// make background
	ctx.SetFillStyle("#ecf0f1")
	ctx.Rectangle(0, 0, 150, 150).Filled().Draw()

	// draw walls
	ctx.SetFillStyle("#2c3e50")
	ctx.Rectangle(12, 12, 150, 150).Rounded(15).Draw()
	ctx.Rectangle(19, 19, 150, 150).Rounded(9).Draw()
	ctx.Rectangle(53, 53, 49, 33).Rounded(10).Draw()
	ctx.Rectangle(53, 119, 49, 16).Rounded(6).Draw()
	ctx.Rectangle(135, 53, 49, 33).Rounded(10).Draw()
	ctx.Rectangle(135, 119, 25, 49).Rounded(10).Draw()

	// draw pacman body
	ctx.SetFillStyle("#f39c12")
	ctx.BeginPath()
	ctx.Arc(37, 37, 13, math.Pi/7, -math.Pi/7)
	ctx.LineTo(31, 37)
	ctx.Fill()

	// draw bread crumbs
	ctx.SetFillStyle("#2c3e50")
	for i := 0; i < 8; i++ {
		ctx.Rectangle(51+i*16, 35, 4, 4).Filled().Draw()
	}
	for i := 0; i < 6; i++ {
		ctx.Rectangle(115, 51+i*16, 4, 4).Filled().Draw()
	}

	for i := 0; i < 8; i++ {
		ctx.Rectangle(51+i*16, 99, 4, 4).Filled().Draw()
	}

	// draw ghost's body
	ctx.BeginPath()
	ctx.MoveTo(83, 116)
	ctx.LineTo(83, 102)
	ctx.BezierCurveTo(83, 94, 89, 88, 97, 88)
	ctx.BezierCurveTo(105, 88, 111, 94, 111, 102)
	ctx.LineTo(111, 116)
	ctx.LineTo(106, 111)
	ctx.LineTo(101, 116)
	ctx.LineTo(97, 111)
	ctx.LineTo(92, 116)
	ctx.LineTo(87, 111)
	ctx.LineTo(83, 116)
	ctx.Fill()

	// draw ghost's eyes
	ctx.SetFillStyle("white")
	ctx.BeginPath()
	ctx.MoveTo(91, 96)
	ctx.BezierCurveTo(88, 96, 87, 99, 87, 101)
	ctx.BezierCurveTo(87, 103, 88, 106, 91, 106)
	ctx.BezierCurveTo(94, 106, 95, 103, 95, 101)
	ctx.BezierCurveTo(95, 99, 94, 96, 91, 96)
	ctx.MoveTo(103, 96)
	ctx.BezierCurveTo(100, 96, 99, 99, 99, 101)
	ctx.BezierCurveTo(99, 103, 100, 106, 103, 106)
	ctx.BezierCurveTo(106, 106, 107, 103, 107, 101)
	ctx.BezierCurveTo(107, 99, 106, 96, 103, 96)
	ctx.Fill()

	// draw ghost's pupils
	ctx.SetFillStyle("black")
	ctx.BeginPath()
	ctx.Arc(101, 102, 2, 0, math.Pi*2)
	ctx.Fill()
	ctx.BeginPath()
	ctx.Arc(89, 102, 2, 0, math.Pi*2)
	ctx.Fill()
}
