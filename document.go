package glowasm

import (
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
	return doc.Get("URL").String()
}

// CharacterSet returns document's encoding.
func (doc *Document) CharacterSet() string {
	return doc.Get("characterSet").String()
}

func (doc *Document) CompatMode() string {
	return doc.Get("compatMode").String()
}

// ContentType returns document's content type.
func (doc *Document) ContentType() string {
	return doc.Get("contentType").String()
}

func (doc *Document) Doctype() string {
	return doc.Get("doctype").String()
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
func (doc *Document) Body() HTMLElement {
	return doc.Get("body").HTMLElement()
}

// HTML returns the Element that is a direct child of the document.
// For HTML documents, this is normally the <html> element.
func (doc *Document) HTML() HTMLElement {
	return doc.Get("documentElement").HTMLElement()
}

// Head returns the <head> element of the current document.
func (doc *Document) Head() HTMLElement {
	return doc.Get("head").HTMLElement()
}

func (doc *Document) Embeds() []HTMLElement {
	collection := doc.Get("embeds")
	values := collection.Values()
	elements := make([]HTMLElement, len(values), 0)
	for i, value := range values {
		elements[i] = value.HTMLElement()
	}
	return elements
}

// NON-STRING PROPERTIES

func (doc *Document) ChildElementCount() int {
	return doc.Get("childElementCount").Int()
}

// DesignMode indicates whether the document can be edited.
func (doc *Document) DesignMode() bool {
	return doc.Get("designMode").String() == "on"
}

func (doc *Document) FullscreenEnabled() bool {
	return doc.Get("fullscreenEnabled").Bool()
}

func (doc *Document) Hidden() bool {
	return doc.Get("hidden").Bool()
}

func (doc *Document) LastModified() time.Time {
	date := doc.Get("lastModified").String()
	timestamp := doc.Get("Date").Call("parse", date).Float()
	return time.Unix(0, int64(timestamp))
}

func (doc *Document) XMLStandalone() bool {
	return doc.Get("xmlStandalone").Bool()
}

// METHODS

func (doc *Document) CreateElement(name string) HTMLElement {
	return doc.Call("createElement", name).HTMLElement()
}

func (doc *Document) CreateCanvas() Canvas {
	return doc.CreateElement("canvas").Canvas()
}

// SUBTYPES

type Fullscreen struct {
	value Value
}

func (scroll *Scroll) Available() bool {
	return scroll.value.Get("fullscreenEnabled").Bool()
}
