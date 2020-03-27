package web

// https://developer.mozilla.org/en-US/docs/Web/API/Screen
type Screen struct {
	Value
}

// PROPERTIES

// Returns the height of the screen, in pixels,
// minus permanent or semipermanent user interface features
// displayed by the operating system, such as the Taskbar on Windows.
// https://developer.mozilla.org/en-US/docs/Web/API/Screen/availHeight
func (screen Screen) AvailableHeight() int {
	return screen.Get("availHeight").Int()
}

// Returns the amount of horizontal space in pixels available to the window.
// https://developer.mozilla.org/en-US/docs/Web/API/Screen/availWidth
func (screen Screen) AvailableWidth() int {
	return screen.Get("availWidth").Int()
}

// Returns the height of the screen in pixels.
// https://developer.mozilla.org/en-US/docs/Web/API/Screen/height
func (screen Screen) Height() int {
	return screen.Get("height").Int()
}

// Returns the width of the screen.
// https://developer.mozilla.org/en-US/docs/Web/API/Screen/width
func (screen Screen) Width() int {
	return screen.Get("width").Int()
}
