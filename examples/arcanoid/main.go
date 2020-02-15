package main

import (
	"fmt"
	"time"

	"github.com/life4/gweb/canvas"
	"github.com/life4/gweb/web"
)

const BGColor = "#ecf0f1"
const PlatformColor = "#2c3e50"
const TextColor = "#2c3e50"

const PlatformWidth = 100
const PlatformHeight = 30
const PlatformMaxSpeed = 40

type Platfrom struct {
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

func (ctx *Platfrom) changePosition() {
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

func (ctx *Platfrom) handleMouse(event web.Event) {
	ctx.mouseX = event.Get("clientX").Int()
}

func (ctx *Platfrom) handleFrame() {
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

	// register mouse movement handler
	platform := Platfrom{
		context:      context,
		element:      canvas,
		x:            w / 2,
		mouseX:       w / 2,
		width:        PlatformWidth,
		windowWidth:  w,
		windowHeight: h,
	}
	canvas.EventTarget().Listen(web.EventTypeMouseMove, platform.handleMouse)

	// register frame updaters
	fps := FPS{context: context, updated: time.Now()}
	handler := func() {
		fps.handle()
		platform.handleFrame()
	}
	window.RequestAnimationFrame(handler, true)

	// prevent ending of the program
	select {}
}
