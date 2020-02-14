package main

import (
	"fmt"

	"github.com/life4/gweb/web"
)

func handleMouseMove(event web.Event) {
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	element := event.CurrentTarget().HTMLElement()
	text := fmt.Sprintf("The mouse position is %d x %d", x, y)
	element.SetText(text)
}

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Events handling example")

	// create <div>
	div := doc.CreateElement("div")
	div.SetText("no movement")

	// fill all the page by the element
	h := window.InnerHeight()
	w := window.InnerWidth()
	div.Style().SetHeight(fmt.Sprintf("%dpx", h), false)
	div.Style().SetWidth(fmt.Sprintf("%dpx", w), false)

	// register the listener
	div.EventTarget().Listen(web.EventTypeMouseMove, handleMouseMove)

	// add the element into <body>
	body := doc.Body()
	body.Node().AppendChild(div.Node())

	// prevent the script from stopping
	select {}
}
