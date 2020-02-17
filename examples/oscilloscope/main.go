package main

import (
	"math"
	"sync"

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
		y := scope.GetY(freq, painter.height)
		painter.context.LineTo(int(math.Round(x)), y)
		x += chunkWidth
	}
	painter.context.LineTo(painter.width, painter.height/2)
	painter.context.Stroke()
}

type ScopeDomain struct {
	data *audio.TimeDomainBytes
}

func (scope *ScopeDomain) Data() []byte {
	scope.data.Update()
	return scope.data.Data
}

func (scope *ScopeDomain) Size() int {
	return scope.data.Size
}

func (scope *ScopeDomain) GetY(value byte, height int) int {
	return height - int(value)*height/256
}

type ScopeFreq struct {
	data *audio.FrequencyDataBytes
}

func (scope *ScopeFreq) Data() []byte {
	scope.data.Update()
	return scope.data.Data
}

func (scope *ScopeFreq) Size() int {
	return scope.data.Size
}

func (scope *ScopeFreq) GetY(value byte, height int) int {
	return height - 10 - int(value)*(height-10)/256
}

func makeCanvas(w, h int) canvas.Context2D {
	window := web.GetWindow()
	doc := window.Document()
	body := doc.Body()

	canvas := doc.CreateCanvas()
	canvas.SetHeight(h)
	canvas.SetWidth(w)
	body.Node().AppendChild(canvas.Node())

	context := canvas.Context2D()

	// draw background
	context.SetFillStyle(BGColor)
	context.Rectangle(0, 0, w, h).Filled().Draw()

	return context
}

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Audio visualization example")

	h := window.InnerHeight()/2 - 40
	w := window.InnerWidth() - 40

	var domain audio.TimeDomainBytes
	var freq audio.FrequencyDataBytes

	go func() {
		// get audio stream from mic
		promise := window.Navigator().MediaDevices().Audio()
		msg, err := promise.Get()
		if err.Truthy() {
			window.Console().Error("", err)
		}
		stream := msg.MediaStream()

		// make analyzer and update time domain and frequency managers
		audioContext := window.AudioContext()
		analyser := audioContext.Analyser()
		analyser.SetMinDecibels(-90)
		analyser.SetMaxDecibels(-10)
		analyser.SetSmoothingTimeConstant(0.85)
		analyser.SetFFTSize(1024)
		domain = analyser.TimeDomain()
		freq = analyser.FrequencyData()

		// connect audio context to the stream
		source := audioContext.MediaStreamSource(stream)
		source.Connect(analyser.AudioNode, 0, 0)
	}()

	// make domain data painting handler
	scopeD := ScopeDomain{
		data: &domain,
	}
	painterD := Painter{
		context: makeCanvas(w, h),
		width:   w,
		height:  h,
	}

	// make frequency data painting handler
	scopeF := ScopeFreq{
		data: &freq,
	}
	painterF := Painter{
		context: makeCanvas(w, h),
		width:   w,
		height:  h,
	}

	// register handlers
	handle := func() {
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			painterD.handle(&scopeD)
			wg.Done()
		}()
		go func() {
			painterF.handle(&scopeF)
			wg.Done()
		}()
		wg.Wait()
	}
	window.RequestAnimationFrame(handle, true)
	// prevent ending of the program
	select {}
}
