package glowasm

import (
	"syscall/js"
	"time"
)

type Document struct {
	value js.Value
}

func GetDocument() Document {
	window := js.Global()
	return Document{value: window.Get("document")}
}

// DOCUMENT STRING PROPERTIES

// URL returns the URL for the current document.
func (doc *Document) URL() string {
	return doc.value.Get("URL").String()
}

// Cookie returns the HTTP cookies that apply to the Document.
// If there are no cookies or cookies can't be applied to this resource, the empty string will be returned.
func (doc *Document) Cookie() string {
	return doc.value.Get("URL").String()
}

// CharacterSet returns document's encoding.
func (doc *Document) CharacterSet() string {
	return doc.value.Get("characterSet").String()
}

func (doc *Document) CompatMode() string {
	return doc.value.Get("compatMode").String()
}

// ContentType returns document's content type.
func (doc *Document) ContentType() string {
	return doc.value.Get("contentType").String()
}

func (doc *Document) Domain() string {
	v := doc.value.Get("domain")
	switch v.Type() {
	case js.TypeUndefined, js.TypeNull:
		return ""
	case js.TypeString:
		return v.String()
	default:
		panic("bad type")
	}
}

func (doc *Document) Referrer() string {
	return doc.value.Get("referrer").String()
}

func (doc *Document) InputEncoding() string {
	return doc.value.Get("inputEncoding").String()
}

func (doc *Document) ReadyState() string {
	return doc.value.Get("readyState").String()
}

func (doc *Document) Title() string {
	return doc.value.Get("title").String()
}

// DOCUMENT NON-STRING PROPERTIES

func (doc *Document) ChildElementCount() int {
	return doc.value.Get("childElementCount").Int()
}

// DesignMode indicates whether the document can be edited.
func (doc *Document) DesignMode() bool {
	return doc.value.Get("designMode").String() == "on"
}

func (doc *Document) FullscreenEnabled() bool {
	return doc.value.Get("fullscreenEnabled").Bool()
}

func (doc *Document) Hidden() bool {
	return doc.value.Get("hidden").Bool()
}

func (doc *Document) LastModified() time.Time {
	date := doc.value.Get("lastModified").String()
	timestamp := doc.value.Get("Date").Call("parse", date).Float()
	return time.Unix(0, int64(timestamp))
}

func (doc *Document) XMLStandalone() bool {
	return doc.value.Get("xmlStandalone").Bool()
}

// JS FUNCS

func (doc *Document) CreateElement(name string) Value {
	value := doc.value.Call("createElement", name)
	return Value{value: value}
}

// HELPER FUNCS

func (doc *Document) CreateCanvas() Canvas {
	value := doc.CreateElement("canvas")
	return value.Canvas()
}
