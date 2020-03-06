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

// CONSTRUCTORS

func (window Window) AudioContext() audio.AudioContext {
	constructor := window.Get("AudioContext")
	value := constructor.New().Value
	return audio.AudioContext{Value: value}
}

// SUBTYPE GETTERS

// https://developer.mozilla.org/en-US/docs/Web/API/Window/console
// https://developer.mozilla.org/en-US/docs/Web/API/Console
func (window Window) Console() Console {
	return Console{Value: window.Get("console")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/document
// https://developer.mozilla.org/en-US/docs/Web/API/Document
func (window Window) Document() Document {
	return Document{Value: window.Get("document")}
}

func (window Window) Event() Event {
	return Event{Value: window.Get("event")}
}

func (window Window) Navigator() Navigator {
	return Navigator{Value: window.Get("navigator")}
}

func (window Window) Screen() Screen {
	return Screen{Value: window.Get("screen")}
}

// OTHER GETTERS

func (window Window) InnerHeight() int {
	return window.Get("innerHeight").Int()
}

func (window Window) InnerWidth() int {
	return window.Get("innerWidth").Int()
}

func (window Window) OuterHeight() int {
	return window.Get("outerHeight").Int()
}

func (window Window) OuterWidth() int {
	return window.Get("outerWidth").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/screenX
func (window Window) ScreenX() int {
	return window.Get("screenX").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/screenY
func (window Window) ScreenY() int {
	return window.Get("screenY").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollX
func (window Window) ScrollX() int {
	return window.Get("scrollX").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollY
func (window Window) ScrollY() int {
	return window.Get("scrollY").Int()
}

// SETTERS

// https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollX
func (window Window) SetScrollX(pixels int) {
	window.Set("scrollX", pixels)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollY
func (window Window) SetScrollY(pixels int) {
	window.Set("scrollY", pixels)
}

// METHODS

// https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollY
func (window Window) RequestAnimationFrame(handler func(), recursive bool) {
	wrapped := func(this js.Value, args []js.Value) interface{} {
		handler()
		if recursive {
			window.RequestAnimationFrame(handler, true)
		}
		return nil
	}
	window.Call("requestAnimationFrame", js.FuncOf(wrapped))
}
