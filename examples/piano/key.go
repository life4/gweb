package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/life4/gweb/web"
)

type Key struct {
	Octave  int
	Note    string
	element web.HTMLElement
}

func (key Key) Press() {
	key.element.Style().SetBackgroundColor("#2980b9", false)
	key.element.Style().SetColor("#ecf0f1", false)
}

func (key Key) Release() {
	if strings.Contains(key.Note, "#") {
		key.element.Style().SetBackgroundColor("#7f8c8d", false)
	} else {
		key.element.Style().SetBackgroundColor("#2c3e50", false)
	}
	key.element.Style().SetColor("#bdc3c7", false)
}

func (key *Key) Render(doc web.Document) web.HTMLElement {
	element := doc.CreateElement("span")
	element.SetText(key.Note)
	element.SetID(fmt.Sprintf("key-%d-%s", key.Octave, strings.ReplaceAll(key.Note, "#", "s")))
	element = StyleBlock(element)

	key.element = element
	key.Release()
	return element
}

func KeyFromElement(element web.HTMLElement) Key {
	parts := strings.Split(element.ID(), "-")
	octave, _ := strconv.Atoi(parts[1])
	note := strings.ReplaceAll(parts[2], "s", "#")
	return Key{
		element: element,
		Octave:  octave,
		Note:    note,
	}
}

func KeyFromNote(doc web.Document, octave int, note string) Key {
	id := fmt.Sprintf("key-%d-%s", octave, strings.ReplaceAll(note, "#", "s"))
	element := doc.Element(id)
	if !element.Truthy() {
		return Key{}
	}
	return Key{
		element: element,
		Octave:  octave,
		Note:    note,
	}
}

func StyleBlock(element web.HTMLElement) web.HTMLElement {
	element.Style().SetDisplay("inline-block", false)
	element.Style().SetWidth("40px", false)
	element.Style().SetTextAlign("center", false)
	element.Style().SetMargin("2px", false)
	return element
}
