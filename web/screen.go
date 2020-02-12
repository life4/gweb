package web

type Screen struct {
	Value
}

// PROPERTIES

func (screen Screen) AvailableHeight() int {
	return screen.Get("availHeight").Int()
}

func (screen Screen) AvailableWidth() int {
	return screen.Get("availWidth").Int()
}

func (screen Screen) Height() int {
	return screen.Get("height").Int()
}

func (screen Screen) Width() int {
	return screen.Get("width").Int()
}
