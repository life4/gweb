package web

import (
	"syscall/js"

	"github.com/life4/gweb/audio"
)

type Window struct {
	Value
}

func GetWindow() Window {
	value := Value{Value: js.Global()}
	return Window{Value: value}
}

func (window Window) AudioContext() audio.AudioContext {
	constructor := window.Get("AudioContext")
	jsV := constructor.New().Value
	audioV := audio.Value{Value: jsV}
	return audio.AudioContext{Value: audioV}
}

func (window Window) Console() Console {
	return Console{Value: window.Get("console")}
}

func (window Window) Document() Document {
	return Document{Value: window.Get("document")}
}

func (window Window) Navigator() Navigator {
	return Navigator{Value: window.Get("navigator")}
}

func (window Window) Screen() Screen {
	return Screen{Value: window.Get("screen")}
}
