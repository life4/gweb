package web

import (
	"syscall/js"
	"time"
)

type Document struct {
	Value
}

// SUBTYPE GETTERS

func (doc *Document) Fullscreen() Fullscreen {
	return Fullscreen{value: doc.Value}
}

func (doc *Document) Node() Node {
	return Node{value: doc.Value}
}

// DOCUMENT STRING PROPERTIES

// URL returns the URL for the current document.
func (doc *Document) URL() string {
	return doc.Get("URL").String()
}

// Cookie returns the HTTP cookies that apply to the Document.
// If there are no cookies or cookies can't be applied to this resource, the empty string will be returned.
func (doc *Document) Cookie() string {
	return doc.Get("cookie").String()
}

// CharacterSet returns document's encoding.
func (doc *Document) CharacterSet() string {
	return doc.Get("characterSet").String()
}

// ContentType returns document's content type.
func (doc *Document) ContentType() string {
	return doc.Get("contentType").String()
}

func (doc *Document) DocType() string {
	return doc.Get("doctype").Get("name").String()
}

func (doc *Document) Domain() string {
	v := doc.Get("domain")
	return v.OptionalString()
}

func (doc *Document) Referrer() string {
	return doc.Get("referrer").String()
}

func (doc *Document) InputEncoding() string {
	return doc.Get("inputEncoding").String()
}

func (doc *Document) ReadyState() string {
	return doc.Get("readyState").String()
}

func (doc *Document) Title() string {
	return doc.Get("title").String()
}

// GETTING CONCRETE SUBELEMENTS

// Body returns the <body> or <frameset> node of the current document.
func (doc Document) Body() HTMLElement {
	return doc.Get("body").HTMLElement()
}

// Head returns the <head> element of the current document.
func (doc Document) Head() HTMLElement {
	return doc.Get("head").HTMLElement()
}

// HTML returns the Element that is a direct child of the document.
// For HTML documents, this is normally the <html> element.
func (doc Document) HTML() HTMLElement {
	return doc.Get("documentElement").HTMLElement()
}

// Embeds returns <object> and <embed> elements in the document.
func (doc *Document) Embeds() []Embed {
	collection := doc.Get("plugins")
	values := collection.Values()

	collection = doc.Get("embeds")
	values = append(values, collection.Values()...)

	elements := make([]Embed, len(values))
	for i, value := range values {
		elements[i] = value.Embed()
	}
	return elements
}

// NON-STRING PROPERTIES

// DesignMode indicates whether the document can be edited.
func (doc *Document) DesignMode() bool {
	return doc.Get("designMode").String() == "on"
}

// Hidden is true when the webpage is in the background and not visible to the user
func (doc *Document) Hidden() bool {
	return doc.Get("hidden").Bool()
}

func (doc *Document) LastModified() time.Time {
	date := doc.Get("lastModified").String()
	timestamp := js.Global().Get("Date").Call("parse", date).Float()
	return time.Unix(int64(timestamp/1000), 0)
}

// METHODS

func (doc Document) CreateElement(name string) HTMLElement {
	return doc.Call("createElement", name).HTMLElement()
}

func (doc Document) CreateCanvas() Canvas {
	return doc.CreateElement("canvas").Canvas()
}

func (doc Document) Element(id string) HTMLElement {
	return doc.Call("getElementById", id).HTMLElement()
}

// SUBTYPES

type Fullscreen struct {
	value Value
}

func (scroll *Scroll) Available() bool {
	return scroll.value.Get("fullscreenEnabled").Bool()
}
