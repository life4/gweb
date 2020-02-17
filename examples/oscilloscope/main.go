package main

import (
	"math"

	"github.com/life4/gweb/audio"

	"github.com/life4/gweb/canvas"
	"github.com/life4/gweb/web"
)

const BGColor = "#2c3e50"
const LineColor = "#2ecc71"

type Scope interface {
	Size() int
	Data() []byte
	GetY(value byte, height int) int
}

type Painter struct {
	context canvas.Context2D
	width   int
	height  int
}

func (painter *Painter) handle(scope Scope) {
	// make background (and remove prev results)
	painter.context.SetFillStyle(BGColor)
	painter.context.BeginPath()
	painter.context.Rectangle(0, 0, painter.width, painter.height).Filled().Draw()
	painter.context.ClosePath()
	painter.context.MoveTo(0, painter.height/2)

	// don't draw the line if TimeDomain hasn't been initialized yet
	if scope.Size() == 0 {
		return
	}

	// draw the line
	chunkWidth := float64(painter.width) / float64(scope.Size())
	painter.context.SetFillStyle(LineColor)
	painter.context.Line().SetWidth(2)
	x := 0.0
	for _, freq := range scope.Data() {
		y := int(freq) * painter.height / 256
		painter.context.LineTo(int(math.Round(x)), y)
		x += chunkWidth
	}
	painter.context.LineTo(painter.width, painter.height/2)
	painter.context.Stroke()
}

type ScopeDomain struct {
	domain *audio.TimeDomainBytes
}

func (scope *ScopeDomain) Data() []byte {
	scope.domain.Update()
	return scope.domain.Data
}

func (scope *ScopeDomain) Size() int {
	return scope.domain.Size
}

func (scope *ScopeDomain) GetY(value byte, height int) int {
	return int(value) * height / 256
}

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Audio visualization example")
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
	context.Rectangle(0, 0, w, h).Filled().Draw()

	var domain audio.TimeDomainBytes

	go func() {
		// get audio stream from mic
		promise := window.Navigator().MediaDevices().Audio()
		msg, err := promise.Get()
		if err.Truthy() {
			window.Console().Error("", err)
		}
		stream := msg.MediaStream()

		// make analyzer and update time domain manager
		audioContext := window.AudioContext()
		analyser := audioContext.Analyser()
		analyser.SetMinDecibels(-90)
		analyser.SetMaxDecibels(-10)
		analyser.SetSmoothingTimeConstant(0.85)
		analyser.SetFFTSize(1024)
		domain = analyser.TimeDomain()

		// connect audio context to the stream
		source := audioContext.MediaStreamSource(stream)
		source.Connect(analyser.AudioNode, 0, 0)
	}()

	// register handlers
	scopeD := ScopeDomain{
		domain: &domain,
	}
	painter := Painter{
		context: context,
		width:   w,
		height:  h,
	}
	handle := func() {
		painter.handle(&scopeD)
	}
	window.RequestAnimationFrame(handle, true)
	// prevent ending of the program
	select {}
}
