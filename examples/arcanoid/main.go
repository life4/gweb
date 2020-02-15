package main

import (
	"fmt"
	"math"
	"time"

	"github.com/life4/gweb/canvas"
	"github.com/life4/gweb/web"
)

const BGColor = "#ecf0f1"
const BallColor = "#d35400"
const PlatformColor = "#2c3e50"
const TextColor = "#2c3e50"

const PlatformWidth = 100
const PlatformHeight = 30
const PlatformMaxSpeed = 40
const BallSize = 30

type Platform struct {
	context canvas.Context2D
	element web.Canvas
	// geometry
	width int
	x     int
	// movement
	mouseX int
	// borders
	windowWidth  int
	windowHeight int
}

func (ctx *Platform) changePosition() {
	path := ctx.mouseX - (ctx.x + ctx.width/2)
	if path == 0 {
		return
	}
	if path > 0 && path > PlatformMaxSpeed {
		path = PlatformMaxSpeed
	} else if path < 0 && path < -PlatformMaxSpeed {
		path = -PlatformMaxSpeed
	}
	ctx.x += path
}

func (ctx *Platform) handleMouse(event web.Event) {
	ctx.mouseX = event.Get("clientX").Int()
}

func (ctx *Platform) handleFrame() {
	y := ctx.windowHeight - 100

	// clear out previous render
	ctx.context.SetFillStyle(BGColor)
	ctx.context.Rectangle(ctx.x, y, ctx.width, PlatformHeight).Filled().Draw()

	// change platform coordinates
	ctx.changePosition()

	// draw the platform
	ctx.context.SetFillStyle(PlatformColor)
	ctx.context.Rectangle(ctx.x, y, ctx.width, PlatformHeight).Filled().Draw()
}

type Ball struct {
	context canvas.Context2D
	// position
	x, y int
	// movement
	vectorX int
	vectorY int
	// borders
	windowWidth  int
	windowHeight int
	platform     *Platform
}

func (ctx *Ball) changeDirection() {
	ballX := ctx.x + ctx.vectorX
	ballY := ctx.y + ctx.vectorY
	// bounce from text box (where we draw FPS and score)
	if ballX < 110+BallSize && ballY < 60 {
		ctx.vectorX = -ctx.vectorX
	}
	if ballX < 110 && ballY < 60+BallSize {
		ctx.vectorY = -ctx.vectorY
	}

	// right and left
	if ballX > ctx.windowWidth-BallSize {
		ctx.vectorX = -ctx.vectorX
	} else if ballX < BallSize {
		ctx.vectorX = -ctx.vectorX
	}

	// bottom and top
	if ballY > ctx.windowHeight-BallSize {
		ctx.vectorY = -ctx.vectorY
	} else if ballY < BallSize {
		ctx.vectorY = -ctx.vectorY
	}

	// bounce from platform top
	if ctx.vectorY > 0 {
		platformTop := ctx.windowHeight - 100
		if ballY+BallSize > platformTop && ballY+BallSize <= platformTop+PlatformHeight {
			platformLeft := ctx.platform.x
			platformRight := ctx.platform.x + ctx.platform.width
			// ball touched the platform by the bottom
			if ballX >= platformLeft && ballX <= platformRight {
				ctx.vectorY = -ctx.vectorY
			}
		}
	}
}

func (ctx *Ball) handle() {
	// clear out previous render
	ctx.context.SetFillStyle(BGColor)
	ctx.context.BeginPath()
	ctx.context.Arc(ctx.x, ctx.y, BallSize+1, 0, math.Pi*2)
	ctx.context.Fill()
	ctx.context.ClosePath()

	ctx.changeDirection()

	// move the ball
	ctx.x += ctx.vectorX
	ctx.y += ctx.vectorY

	// draw the ball
	ctx.context.SetFillStyle(BallColor)
	ctx.context.BeginPath()
	ctx.context.Arc(ctx.x, ctx.y, BallSize, 0, math.Pi*2)
	ctx.context.Fill()
	ctx.context.ClosePath()
}

type FPS struct {
	context canvas.Context2D
	updated time.Time
}

func (h *FPS) drawFPS(now time.Time) {
	// calculate FPS
	fps := time.Second / now.Sub(h.updated)
	text := fmt.Sprintf("%d FPS", int64(fps))

	// clear
	h.context.SetFillStyle(BGColor)
	h.context.Rectangle(10, 10, 100, 20).Filled().Draw()

	// write
	h.context.Text().SetFont("bold 20px Roboto")
	h.context.SetFillStyle(TextColor)
	h.context.Text().Fill(text, 10, 30, 100)
}

func (h *FPS) handle() {
	now := time.Now()
	// update FPS counter every second
	if h.updated.Second() != now.Second() {
		h.drawFPS(now)
	}
	h.updated = now
}

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Canvas drawing example")
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
		context:      context,
		element:      canvas,
		x:            w / 2,
		mouseX:       w / 2,
		width:        PlatformWidth,
		windowWidth:  w,
		windowHeight: h,
	}
	fps := FPS{context: context, updated: time.Now()}
	ball := Ball{
		context: context,
		vectorX: 4, vectorY: -4,
		x: 120, y: 120,
		windowWidth: w, windowHeight: h,
		platform: &platform,
	}

	// register mouse movement handler
	canvas.EventTarget().Listen(web.EventTypeMouseMove, platform.handleMouse)

	// register frame updaters
	handler := func() {
		go fps.handle()
		ball.handle()
		platform.handleFrame()
	}
	window.RequestAnimationFrame(handler, true)

	// prevent ending of the program
	select {}
}
