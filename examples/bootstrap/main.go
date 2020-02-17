package main

import (
	"sync"

	"github.com/life4/gweb/web"
)

type Listener struct {
	sync.WaitGroup
}

func (listener *Listener) showAlert(event web.Event) {
	window := web.GetWindow()
	doc := window.Document()

	// create alert
	div := doc.CreateElement("div")
	div.SetText("It works!")
	div.Class().Append("alert", "alert-success")
	div.Set("role", "alert")

	// add the element into <body>
	body := doc.Body()
	body.Node().AppendChild(div.Node())

	// allow to close the program (unblock `listener.Wait()`)
	listener.Done()
}

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Twitter bootstrap including example")

	// make <link/>
	link := doc.CreateElement("link")
	// since we have to set element-specific fields, we have to go in syscall/js style
	link.Set("rel", "stylesheet")
	link.Set("href", "https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css")

	// register listener for <link/> load to show message only when CSS is ready
	listener := Listener{}
	listener.Add(1)
	link.EventTarget().Listen(web.EventTypeLoad, listener.showAlert)

	// add <link/> into <head>
	head := doc.Head()
	head.Node().AppendChild(link.Node())

	// wait for listener to end before closing the program
	listener.Wait()
}
