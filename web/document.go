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
// https://developer.mozilla.org/en-US/docs/Web/API/Document/URL
func (doc *Document) URL() string {
	return doc.Get("URL").String()
}

// Cookie returns the HTTP cookies that apply to the Document.
// If there are no cookies or cookies can't be applied to this resource, the empty string will be returned.
func (doc *Document) Cookie() string {
	return doc.Get("cookie").String()
}

// CharacterSet returns document's encoding.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/characterSet
func (doc *Document) CharacterSet() string {
	return doc.Get("characterSet").String()
}

// ContentType returns document's content type.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/contentType
func (doc *Document) ContentType() string {
	return doc.Get("contentType").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/doctype
func (doc *Document) DocType() string {
	return doc.Get("doctype").Get("name").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/domain
func (doc *Document) Domain() string {
	v := doc.Get("domain")
	return v.OptionalString()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/referrer
func (doc *Document) Referrer() string {
	return doc.Get("referrer").String()
}

func (doc *Document) InputEncoding() string {
	return doc.Get("inputEncoding").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/readyState
func (doc *Document) ReadyState() string {
	return doc.Get("readyState").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/title
func (doc *Document) Title() string {
	return doc.Get("title").String()
}

// GETTING CONCRETE SUBELEMENTS

// Body returns the <body> or <frameset> node of the current document.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (doc Document) Body() HTMLElement {
	return doc.Get("body").HTMLElement()
}

// Head returns the <head> element of the current document.
func (doc Document) Head() HTMLElement {
	return doc.Get("head").HTMLElement()
}

// HTML returns the Element that is a direct child of the document.
// For HTML documents, this is normally the <html> element.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/head
func (doc Document) HTML() HTMLElement {
	return doc.Get("documentElement").HTMLElement()
}

// Embeds returns <object> and <embed> elements in the document.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/embeds
// https://developer.mozilla.org/en-US/docs/Web/API/Document/plugins
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
// https://developer.mozilla.org/en-US/docs/Web/API/Document/designMode
func (doc *Document) DesignMode() bool {
	return doc.Get("designMode").String() == "on"
}

// Hidden is true when the webpage is in the background and not visible to the user
// https://developer.mozilla.org/en-US/docs/Web/API/Document/hidden
func (doc *Document) Hidden() bool {
	return doc.Get("hidden").Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document/lastModified
func (doc *Document) LastModified() time.Time {
	date := doc.Get("lastModified").String()
	timestamp := js.Global().Get("Date").Call("parse", date).Float()
	return time.Unix(int64(timestamp/1000), 0)
}

// SETTERS

func (doc Document) SetTitle(title string) {
	doc.Set("title", title)
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
