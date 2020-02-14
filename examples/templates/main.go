package main

import "github.com/life4/gweb/web"

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Templates example")
	body := doc.Body()

	// create template
	template := doc.CreateElement("div")
	template.Style().SetBorder("solid 4px red", false)

	// add <slot> into template
	slot := doc.CreateElement("slot")
	slot.Set("name", "example") // here we call syscall/js-like method
	slot.SetInnerHTML("default text")
	template.Node().AppendChild(slot.Node())

	// Add template into a shadow DOM.
	// This is the most important thing to make the template renderable
	shadow := body.Shadow().Attach()
	// since we clone the template, we should add all <slot>'s before it.
	shadow.Node().AppendChild(template.Node().Clone(true))

	// make <span> element that will replace the <slot>
	span := doc.CreateElement("span")
	span.SetText("The template is rendered!")
	span.SetSlot("example")

	// add <span> into <body>, and it will automatically fill <slot>
	body.Node().AppendChild(span.Node())
}
