package main

import (
	"github.com/life4/gweb/audio"
	"github.com/life4/gweb/web"
)

func main() {
	window := web.GetWindow()
	doc := window.Document()
	body := doc.Body()

	audioContext := window.AudioContext()
	if audioContext.State() != audio.AudioContextStateRunning {
		audioContext.Resume()
	}
	dest := audioContext.Destination()
	gain := audioContext.Gain()
	gain.Connect(dest.AudioNode, 0, 0)
	gain.Gain().Set(1.0)

	keyboard := KeyBoard{
		notes:   getNotes(),
		context: &audioContext,
		gain:    &gain,
		oscs:    make(map[int]map[string]*audio.OscillatorNode),
		octave:  3,
	}
	element := keyboard.Render(doc)
	body.Node().AppendChild(element.Node())

	select {}
}
