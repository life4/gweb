package main

import (
	"math"

	"github.com/life4/gweb/audio"

	"github.com/life4/gweb/canvas"
	"github.com/life4/gweb/web"
)

const BGColor = "#2c3e50"
const LineColor = "#2ecc71"

type Scope struct {
	domain  *audio.TimeDomainBytes
	context canvas.Context2D
	width   int
	height  int
}

func (scope *Scope) handle() {
	// make background (and remove prev results)
	scope.context.SetFillStyle(BGColor)
	scope.context.BeginPath()
	scope.context.Rectangle(0, 0, scope.width, scope.height).Filled().Draw()
	scope.context.ClosePath()
	scope.context.MoveTo(0, scope.height/2)

	// don't draw the line if TimeDomain hasn't been initialized yet
	if scope.domain.Size == 0 {
		return
	}

	// draw the line
	chunkWidth := float64(scope.width) / float64(scope.domain.Size)
	scope.context.SetFillStyle(LineColor)
	scope.domain.Update()
	x := 0.0
	for _, freq := range scope.domain.Data {
		y := int(freq) * scope.height / 256
		scope.context.LineTo(int(math.Round(x)), y)
		x += chunkWidth
	}
	scope.context.LineTo(scope.width, scope.height/2)
	scope.context.Stroke()
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
	scope := Scope{
		domain:  &domain,
		context: context,
		width:   w,
		height:  h,
	}
	window.RequestAnimationFrame(scope.handle, true)
	// prevent ending of the program
	select {}
}
