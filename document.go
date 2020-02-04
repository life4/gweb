package glowasm

import (
	"syscall/js"
	"time"
)

type Document struct {
	Value
}

func GetDocument() Document {
	window := js.Global()
	value := Value{Value: window.Get("document")}
	return Document{Value: value}
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

// DOCUMENT NON-STRING PROPERTIES

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

// JS FUNCS

func (doc *Document) CreateElement(name string) Value {
	return doc.Call("createElement", name)
}

// HELPER FUNCS

func (doc *Document) CreateCanvas() Canvas {
	value := doc.CreateElement("canvas")
	return value.Canvas()
}

func (doc *Document) Node() Node {
	return Node{value: doc.Value}
}
