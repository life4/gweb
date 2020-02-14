package main

import (
	"math"

	"github.com/life4/gweb/web"
)

func roundedRect(ctx web.Context2D, x, y, width, height, radius int) {
	ctx.BeginPath()
	ctx.MoveTo(x, y+radius)
	ctx.LineTo(x, y+height-radius)
	ctx.ArcTo(x, y+height, x+radius, y+height, radius)
	ctx.LineTo(x+width-radius, y+height)
	ctx.ArcTo(x+width, y+height, x+width, y+height-radius, radius)
	ctx.LineTo(x+width, y+radius)
	ctx.ArcTo(x+width, y, x+width-radius, y, radius)
	ctx.LineTo(x+radius, y)
	ctx.ArcTo(x, y, x, y+radius, radius)
	ctx.Stroke()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Canvas_API/Tutorial/Drawing_shapes#Making_combinations
func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Canvas triangle example")
	body := doc.Body()

	// create canvas
	canvas := doc.CreateCanvas()
	canvas.SetHeight(150)
	canvas.SetWidth(150)
	body.Node().AppendChild(canvas.Node())

	ctx := canvas.Context2D()

	// make background
	ctx.SetFillStyle("#ecf0f1")
	ctx.FillRect(0, 0, 150, 150)

	// draw walls
	ctx.SetFillStyle("#2c3e50")
	roundedRect(ctx, 12, 12, 150, 150, 15)
	roundedRect(ctx, 19, 19, 150, 150, 9)
	roundedRect(ctx, 53, 53, 49, 33, 10)
	roundedRect(ctx, 53, 119, 49, 16, 6)
	roundedRect(ctx, 135, 53, 49, 33, 10)
	roundedRect(ctx, 135, 119, 25, 49, 10)

	// draw pacman body
	ctx.SetFillStyle("#f39c12")
	ctx.BeginPath()
	ctx.Arc(37, 37, 13, math.Pi/7, -math.Pi/7)
	ctx.LineTo(31, 37)
	ctx.Fill()

	// draw bread crumbs
	ctx.SetFillStyle("#2c3e50")
	for i := 0; i < 8; i++ {
		ctx.FillRect(51+i*16, 35, 4, 4)
	}
	for i := 0; i < 6; i++ {
		ctx.FillRect(115, 51+i*16, 4, 4)
	}

	for i := 0; i < 8; i++ {
		ctx.FillRect(51+i*16, 99, 4, 4)
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
