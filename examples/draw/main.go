package main

import (
	"fmt"
	"math"
	"time"

	"github.com/life4/gweb/web"
)

const BGColor = "#ecf0f1"
const PointColor = "#2c3e50"
const TextColor = "#2c3e50"

type Handler struct {
	context web.Context2D
	drawing bool
	updated time.Time
}

func (h *Handler) handleStart(event web.Event) {
	h.drawing = true
}

func (h *Handler) handleEnd(event web.Event) {
	h.drawing = false
}

func (h *Handler) handleMove(event web.Event) {
	if !h.drawing {
		return
	}

	// draw a point
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()
	h.context.SetFillStyle(PointColor)
	h.context.BeginPath()
	h.context.Arc(x, y, 10, 0, math.Pi*2)
	h.context.Fill()
}

func (h *Handler) handleFrame() {
	now := time.Now()
	// update FPS counter every second
	if h.updated.Second() != now.Second() {
		// calculate FPS
		fps := time.Second / now.Sub(h.updated)
		text := fmt.Sprintf("%d FPS", int64(fps))

		// clear
		h.context.SetFillStyle(BGColor)
		h.context.FillRect(10, 10, 100, 20)

		// write
		h.context.Text().SetFont("bold 20px Roboto")
		h.context.SetFillStyle(TextColor)
		h.context.Text().Fill(text, 10, 30, 100)
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
	context.Rect(0, 0, w, h)
	context.Fill()
	context.ClosePath()

	// register handlers
	handler := Handler{context: context, drawing: false, updated: time.Now()}
	canvas.EventTarget().Listen(web.EventTypeMouseDown, handler.handleStart)
	canvas.EventTarget().Listen(web.EventTypeMouseUp, handler.handleEnd)
	canvas.EventTarget().Listen(web.EventTypeMouseMove, handler.handleMove)

	window.RequestAnimationFrame(handler.handleFrame, true)
	// prevent ending of the program
	select {}
}
