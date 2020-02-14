package main

import "github.com/life4/gweb/web"

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Welcome page")

	// create <h1>
	header := doc.CreateElement("h1")
	header.SetText("Hello!")

	// add the element into <body>
	body := doc.Body()
	body.Node().AppendChild(header.Node())
}
