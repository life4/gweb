package main

import (
	"github.com/life4/gweb/audio"
)

type Sound struct {
	gain    audio.GainNode
	osc     audio.OscillatorNode
	context audio.AudioContext
}

func Play(context audio.AudioContext, freq float64) Sound {
	dest := context.Destination()
	gain := context.Gain()
	gain.Connect(dest.AudioNode, 0, 0)
	gain.Gain().Set(1.0)

	osc := context.Oscillator()
	osc.Connect(gain.AudioNode, 0, 0)
	osc.SetShape(audio.ShapeTriangle)
	osc.Frequency().Set(freq)
	osc.Start(0)

	return Sound{
		osc:     osc,
		context: context,
	}
}

func (sound *Sound) Stop() {
	sound.osc.Stop(0)
}
