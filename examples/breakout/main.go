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

	restartButton := doc.CreateElement("button")
	restartButton.SetText("restart")
	restartHandler := func(event web.Event) {
		go func() {
			game.Stop()
			game.Init()
			game.Register()
		}()
	}
	restartButton.EventTarget().Listen(web.EventTypeMouseDown, restartHandler)
	body.Node().AppendChild(restartButton.Node())

	pauseButton := doc.CreateElement("button")
	pauseButton.SetText("pause")
	pauseHandler := func(event web.Event) {
		go func() {
			if !game.state.Stop.Requested {
				game.Stop()
				pauseButton.SetText("play")
			} else {
				game.Register()
				pauseButton.SetText("pause")
			}
		}()
	}
	pauseButton.EventTarget().Listen(web.EventTypeMouseDown, pauseHandler)
	body.Node().AppendChild(pauseButton.Node())

	// prevent ending of the program
	select {}
}
