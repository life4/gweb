package main

import (
	"fmt"
	"math"
	"time"

	"github.com/life4/gweb/canvas"
	"github.com/life4/gweb/web"
)

const BGColor = "#ecf0f1"
const BallColor = "#2c3e50"
const TextColor = "#2c3e50"

type Ball struct {
	context canvas.Context2D
	size    int
	// position
	x, y int
	// movement
	vectorX int
	vectorY int
	// borders
	windowWidth  int
	windowHeight int
}

func (ctx *Ball) changeDirection() {
	// bounce from text box (where we draw FPS and score)
	if ctx.x+ctx.vectorX < 110+ctx.size && ctx.y+ctx.vectorY < 60 {
		ctx.vectorX = -ctx.vectorX
	}
	if ctx.x+ctx.vectorX < 110 && ctx.y+ctx.vectorY < 60+ctx.size {
		ctx.vectorY = -ctx.vectorY
	}

	// right and left
	if ctx.x+ctx.vectorX > ctx.windowWidth-ctx.size {
		ctx.vectorX = -ctx.vectorX
	} else if ctx.x+ctx.vectorX < ctx.size {
		ctx.vectorX = -ctx.vectorX
	}

	// bottom and top
	if ctx.y+ctx.vectorY > ctx.windowHeight-ctx.size {
		ctx.vectorY = -ctx.vectorY
	} else if ctx.y+ctx.vectorY < ctx.size {
		ctx.vectorY = -ctx.vectorY
	}
}

func (ctx *Ball) handle() {
	ctx.changeDirection()

	// clear out previous render
	ctx.context.SetFillStyle(BGColor)
	ctx.context.Rectangle(ctx.x-ctx.size, ctx.y-ctx.size, ctx.size*2, ctx.size*2).Filled().Draw()

	// move the ball
	ctx.x += ctx.vectorX
	ctx.y += ctx.vectorY

	// draw the ball
	ctx.context.SetFillStyle(BallColor)
	ctx.context.BeginPath()
	ctx.context.Arc(ctx.x, ctx.y, ctx.size, 0, math.Pi*2)
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

type Click struct {
	context canvas.Context2D
	ball    *Ball
	score   int
}

func (ctx *Click) touched() {
	ctx.score += 1

	// speed up
	if ctx.ball.vectorX > 0 {
		ctx.ball.vectorX += 1
	} else {
		ctx.ball.vectorX -= 1
	}
	if ctx.ball.vectorY > 0 {
		ctx.ball.vectorY += 1
	} else {
		ctx.ball.vectorY -= 1
	}

	// change direction
	ctx.ball.vectorX = -ctx.ball.vectorX
	ctx.ball.vectorY = -ctx.ball.vectorY

	// make text
	var text string
	if ctx.score == 1 {
		text = fmt.Sprintf("%d hit", ctx.score)
	} else {
		text = fmt.Sprintf("%d hits", ctx.score)
	}

	// clear place where previous score was
	ctx.context.SetFillStyle(BGColor)
	ctx.context.Rectangle(10, 40, 100, 20).Filled().Draw()

	// draw the score
	ctx.context.SetFillStyle(TextColor)
	ctx.context.Text().SetFont("bold 20px Roboto")
	ctx.context.Text().Fill(text, 10, 60, 100)
}

func (ctx *Click) handle(event web.Event) {
	mouseX := event.Get("clientX").Int()
	mouseY := event.Get("clientY").Int()

	hypotenuse := math.Pow(float64(ctx.ball.size+15), 2)
	cathetus1 := math.Pow(float64(mouseX-ctx.ball.x), 2)
	cathetus2 := math.Pow(float64(mouseY-ctx.ball.y), 2)
	if cathetus1+cathetus2 < hypotenuse {
		go ctx.touched()
	}
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

	// register animation handlers
	ball := Ball{
		context: context,
		vectorX: 4, vectorY: -4,
		size: 35, x: 120, y: 120,
		windowWidth: w, windowHeight: h,
	}
	fps := FPS{context: context, updated: time.Now()}
	handler := func() {
		ball.handle()
		fps.handle()
	}
	window.RequestAnimationFrame(handler, true)

	// register action handlers
	click := Click{context: context, ball: &ball, score: 0}
	canvas.EventTarget().Listen(web.EventTypeMouseDown, click.handle)

	// prevent ending of the program
	select {}
}
