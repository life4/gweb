package web

import (
	"syscall/js"

	"github.com/life4/gweb/audio"
)

// https://developer.mozilla.org/en-US/docs/Web/API/Window
type Window struct {
	Value
}

// Returns JS global
// https://developer.mozilla.org/en-US/docs/Web/API/Window
func GetWindow() Window {
	value := Value{Value: js.Global()}
	return Window{Value: value}
}

// CONSTRUCTORS

// https://developer.mozilla.org/en-US/docs/Web/API/AudioContext
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

// Event returns the current event.
// The Event object passed directly to event handlers should be used instead whenever possible.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/event
func (window Window) Event() Event {
	return Event{Value: window.Get("event")}
}

// Navigator returns a reference to the Navigator object,
// which has methods and properties about the application running the script.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/navigator
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator
func (window Window) Navigator() Navigator {
	return Navigator{Value: window.Get("navigator")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/screen
// https://developer.mozilla.org/en-US/docs/Web/API/Screen
func (window Window) Screen() Screen {
	return Screen{Value: window.Get("screen")}
}

// Create an object used to send HTTP requests (XMLHttpRequest),
// open it (initialize) and set the response type (responseType) to binary.
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/XMLHttpRequest
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/open
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/responseType
func (window Window) HTTPRequest(method, url string) HTTPRequest {
	req := HTTPRequest{Value: window.Get("XMLHttpRequest").New()}
	req.Call("open", method, url, true)
	req.Set("responseType", "blob")
	return req
}

// OTHER GETTERS

// Returns the height of the content area of the browser window including, if rendered, the horizontal scrollbar.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/innerHeight
func (window Window) InnerHeight() int {
	return window.Get("innerHeight").Int()
}

// Returns the width of the content area of the browser window including, if rendered, the vertical scrollbar.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/innerWidth
func (window Window) InnerWidth() int {
	return window.Get("innerWidth").Int()
}

// Returns the height of the outside of the browser window.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/outerHeight
func (window Window) OuterHeight() int {
	return window.Get("outerHeight").Int()
}

// Returns the width of the outside of the browser window.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/outerWidth
func (window Window) OuterWidth() int {
	return window.Get("outerWidth").Int()
}

// Returns the horizontal distance from the left border of the user's browser viewport to the left side of the screen.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/screenX
func (window Window) ScreenX() int {
	return window.Get("screenX").Int()
}

// Returns the vertical distance from the top border of the user's browser viewport to the top side of the screen.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/screenY
func (window Window) ScreenY() int {
	return window.Get("screenY").Int()
}

// Returns the number of pixels that the document has already been scrolled horizontally.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollX
func (window Window) ScrollX() int {
	return window.Get("scrollX").Int()
}

// Returns the number of pixels that the document has already been scrolled vertically.
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
