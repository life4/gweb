package main

import (
	"github.com/life4/gweb/web"
)

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Breakout")
	body := doc.Body()

	// create canvas
	h := window.InnerHeight() - 40
	w := window.InnerWidth() - 40
	canvas := doc.CreateCanvas()
	canvas.SetHeight(h)
	canvas.SetWidth(w)
	body.Node().AppendChild(canvas.Node())

	game := Game{
		Width:  w,
		Height: h,
		Window: window,
		Canvas: canvas,
		Body:   body,
	}
	game.Init()
	game.Register()

	// prevent ending of the program
	select {}
}
