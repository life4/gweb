package main

import (
	"fmt"
	"strings"

	"github.com/life4/gweb/web"
)

type Key struct {
	Octave  int
	Note    string
	element web.HTMLElement
}

func (key Key) Press() {
	key.element.Style().SetBackgroundColor("#34495e", false)
	key.element.Style().SetColor("#ecf0f1", false)
}

func (key Key) Release() {
	key.element.Style().SetBackgroundColor("#2c3e50", false)
	key.element.Style().SetColor("#bdc3c7", false)
}

func (key *Key) Render(doc web.Document) web.HTMLElement {
	element := doc.CreateElement("span")
	element.SetText(key.Note)
	element.SetID(fmt.Sprintf("key-%d-%s", key.Octave, strings.ReplaceAll(key.Note, "#", "s")))

	element.Style().SetDisplay("inline-block", false)
	element.Style().SetWidth("40px", false)
	element.Style().SetTextAlign("center", false)
	element.Style().SetMargin("2px", false)

	key.element = element
	key.Release()
	return element
}
