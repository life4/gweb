package main

import (
	"github.com/life4/gweb/audio"

	"github.com/life4/gweb/canvas"
	"github.com/life4/gweb/web"
)

const BGColor = "#ecf0f1"
const LineColor = "#2c3e50"

type Scope struct {
	domain  audio.TimeDomainBytes
	context canvas.Context2D
	width   int
	height  int
}

func (scope *Scope) handle() {
	chunkWidth := scope.width / scope.domain.Size

	scope.context.SetFillStyle(BGColor)
	scope.context.Rectangle(0, 0, scope.width, scope.height).Filled().Draw()
	scope.context.MoveTo(0, scope.height/2)

	scope.context.SetFillStyle(LineColor)
	scope.domain.Update()
	x := 0
	for _, freq := range scope.domain.Data {
		y := int(freq) * scope.height / 256
		scope.context.LineTo(x, y)
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

	// get audio objects
	audio := window.AudioContext()
	analyser := audio.Analyser()
	domain := analyser.TimeDomain()

	go func() {
		promise := window.Navigator().MediaDevices().Audio()
		msg, err := promise.Get()
		if err.Truthy() {
			window.Console().Error("", err)
		}
		stream := msg.MediaStream()
		audio.MediaStreamSource(stream)
		analyser.Connect(audio.Destination(), 0, 0)
	}()

	// register handlers
	scope := Scope{
		domain:  domain,
		context: context,
		width:   w,
		height:  h,
	}
	window.RequestAnimationFrame(scope.handle, true)
	// prevent ending of the program
	select {}
}
