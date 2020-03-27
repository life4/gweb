package main

import (
	"github.com/life4/gweb/web"
)

func main() {
	window := web.GetWindow()
	doc := window.Document()
	body := doc.Body()
	// audioContext := window.AudioContext()

	keyboard := KeyBoard{
		notes: getNotes(),
		// context: audioContext,
	}
	element := keyboard.Render(doc)
	body.Node().AppendChild(element.Node())
}
