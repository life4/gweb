package main

import (
	"sync"
	"time"

	"github.com/life4/gweb/web"
)

const (
	BGColor       = "#ecf0f1"
	BallColor     = "#27ae60"
	PlatformColor = "#2c3e50"
	TextColor     = "#2c3e50"
)

const PlatformWidth = 120
const PlatformHeight = 20
const PlatformMaxSpeed = 40

const BallSize = 20

const (
	BrickHeight     = 20
	BrickRows       = 8
	BrickCols       = 14
	BrickMarginLeft = 120 // pixels
	BrickMarginTop  = 10  // pixels
	BrickMarginX    = 5   // pixels
	BrickMarginY    = 5   // pixels
)

const (
	TextWidth  = 90
	TextHeight = 20
	TextLeft   = 10
	TextTop    = 10
	TextMargin = 5

	TextBottom = TextTop + (TextHeight+TextMargin)*3
	TextRight  = TextLeft + TextWidth
)

func sign(n float64) float64 {
	if n >= 0 {
		return 1
	}
	return -1
}

type Point struct{ x, y int }

type Rectangle struct{ x, y, width, height int }

func (rectangle Rectangle) Contains(point Point) bool {
	if point.y < rectangle.y { // point upper
		return false
	}
	if point.y > rectangle.y+rectangle.height { // point downer
		return false
	}
	if point.x > rectangle.x+rectangle.width { // point righter
		return false
	}
	if point.x < rectangle.x { // point lefter
		return false
	}
	return true

}

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Breakout")
	body := doc.Body()

	// create canvas
	h := window.InnerHeight() - 40
	w := window.InnerWidth() - 40
	canvas := doc.CreateCanvas()
	canvas.SetHeight(h)
	canvas.SetWidth(w)
	body.Node().AppendChild(canvas.Node())

	context := canvas.Context2D()

	// draw background
	context.SetFillStyle(BGColor)
	context.BeginPath()
	context.Rectangle(0, 0, w, h).Filled().Draw()
	context.Fill()
	context.ClosePath()

	// make handlers
	platform := Platform{
		Rectangle: Rectangle{
			x:      w / 2,
			y:      h - 60,
			width:  PlatformWidth,
			height: PlatformHeight,
		},
		context:      context,
		element:      canvas,
		mouseX:       w / 2,
		windowWidth:  w,
		windowHeight: h,
	}
	block := TextBlock{context: context, updated: time.Now()}
	ball := Ball{
		context:     context,
		vector:      Vector{x: 5, y: -5},
		Point:       Point{x: platform.x, y: platform.y - BallSize},
		radius:      BallSize,
		windowWidth: w, windowHeight: h,
		platform: &platform,
	}
	bricks := Bricks{
		context:      context,
		windowWidth:  w,
		windowHeight: h,
		ready:        false,
		text:         &block,
	}
	go bricks.Draw()

	// register mouse movement handler
	body.EventTarget().Listen(web.EventTypeMouseMove, platform.handleMouse)

	// register frame updaters
	handler := func() {
		wg := sync.WaitGroup{}
		wg.Add(4)
		go func() {
			// update FPS
			block.handle()
			wg.Done()
		}()
		go func() {
			// update platform position
			platform.handleFrame()
			wg.Done()
		}()
		go func() {
			// check if the ball should bounce from a brick
			bricks.Handle(&ball)
			wg.Done()
		}()
		go func() {
			// check if the ball should bounce from border or platform
			ball.handle()
			wg.Done()
		}()
		wg.Wait()
	}
	window.RequestAnimationFrame(handler, true)

	// prevent ending of the program
	select {}
}
