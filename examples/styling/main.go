package main

import "github.com/life4/gweb/web"

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Welcome page")

	// create <p>
	paragraph := doc.CreateElement("p")
	paragraph.SetText("Styled!")

	// make it cool
	style := paragraph.Style()
	style.SetColor("purple", false)
	style.SetFontFamily("Comic Sans MS", false)
	style.SetFontSize("2em", false)

	// add the element into <body>
	body := doc.Body()
	body.Node().AppendChild(paragraph.Node())
}
