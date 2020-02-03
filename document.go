package glowasm

import (
	"syscall/js"
	"time"
)

type Document struct {
	window js.Value
}

func GetDocument() Document {
	return Document{window: js.Global()}
}

// DOCUMENT STRING PROPERTIES

// URL returns the URL for the current document.
func (doc *Document) URL() string {
	return doc.window.Get("URL").String()
}

// Cookie returns the HTTP cookies that apply to the Document.
// If there are no cookies or cookies can't be applied to this resource, the empty string will be returned.
func (doc *Document) Cookie() string {
	return doc.window.Get("URL").String()
}

// CharacterSet returns document's encoding.
func (doc *Document) CharacterSet() string {
	return doc.window.Get("characterSet").String()
}

func (doc *Document) CompatMode() string {
	return doc.window.Get("compatMode").String()
}

// ContentType returns document's content type.
func (doc *Document) ContentType() string {
	return doc.window.Get("contentType").String()
}

func (doc *Document) Domain() string {
	v := doc.window.Get("domain")
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
	return doc.window.Get("referrer").String()
}

func (doc *Document) InputEncoding() string {
	return doc.window.Get("inputEncoding").String()
}

func (doc *Document) ReadyState() string {
	return doc.window.Get("readyState").String()
}

func (doc *Document) Title() string {
	return doc.window.Get("title").String()
}

// DOCUMENT NON-STRING PROPERTIES

func (doc *Document) ChildElementCount() int {
	return doc.window.Get("childElementCount").Int()
}

// DesignMode indicates whether the document can be edited.
func (doc *Document) DesignMode() bool {
	return doc.window.Get("designMode").String() == "on"
}

func (doc *Document) FullscreenEnabled() bool {
	return doc.window.Get("fullscreenEnabled").Bool()
}

func (doc *Document) Hidden() bool {
	return doc.window.Get("hidden").Bool()
}

func (doc *Document) LastModified() time.Time {
	date := doc.window.Get("lastModified").String()
	timestamp := doc.window.Get("Date").Call("parse", date).Float()
	return time.Unix(0, int64(timestamp))
}

func (doc *Document) XMLStandalone() bool {
	return doc.window.Get("xmlStandalone").Bool()
}
