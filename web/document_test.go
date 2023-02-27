package web

import (
	"strings"
	"syscall/js"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDocumentURL(t *testing.T) {
	d := GetWindow().Document()
	assert.True(t, strings.HasPrefix(d.URL(), "http://127.0.0.1:"), "bad URL")
}

func TestDocumentCookie(t *testing.T) {
	d := GetWindow().Document()
	assert.Equal(t, d.Cookie(), "", "bad cookie string")
}

func TestDocumentCharacterSet(t *testing.T) {
	d := GetWindow().Document()
	assert.Equal(t, d.CharacterSet(), "UTF-8")
}

func TestDocumentContentType(t *testing.T) {
	d := GetWindow().Document()
	assert.Equal(t, d.ContentType(), "text/html")
}

func TestDocumentDocType(t *testing.T) {
	d := GetWindow().Document()
	assert.Equal(t, d.DocType(), "html")
}

func TestDocumentDomain(t *testing.T) {
	d := GetWindow().Document()
	assert.Equal(t, d.Domain(), "127.0.0.1")
}

func TestDocumentReferrer(t *testing.T) {
	d := GetWindow().Document()
	assert.Equal(t, d.Referrer(), "")
}

func TestDocumentReadyState(t *testing.T) {
	d := GetWindow().Document()
	assert.Equal(t, d.ReadyState(), "complete")
}

func TestDocumentTitle(t *testing.T) {
	d := GetWindow().Document()
	assert.Equal(t, d.Title(), "Go wasm")
}

func TestDocumentBody(t *testing.T) {
	d := GetWindow().Document()
	element := d.Body()
	assert.Equal(t, element.Type(), js.TypeObject)
	assert.Equal(t, element.Call("toString").String(), "[object HTMLBodyElement]")
}

func TestDocumentHead(t *testing.T) {
	d := GetWindow().Document()
	element := d.Head()
	assert.Equal(t, element.Type(), js.TypeObject)
	assert.Equal(t, element.Call("toString").String(), "[object HTMLHeadElement]")
}

func TestDocumentHTML(t *testing.T) {
	d := GetWindow().Document()
	element := d.HTML()
	assert.Equal(t, element.Type(), js.TypeObject)
	assert.Equal(t, element.Call("toString").String(), "[object HTMLHtmlElement]")
}

func TestDocumentDesignMode(t *testing.T) {
	d := GetWindow().Document()
	assert.False(t, d.DesignMode())
}

func TestDocumentHidden(t *testing.T) {
	d := GetWindow().Document()
	assert.False(t, d.Hidden())
}

func TestDocumentLastModified(t *testing.T) {
	d := GetWindow().Document()
	assert.WithinDuration(t, d.LastModified(), time.Now(), 5*time.Second)
}

func TestDocumentCreateNode(t *testing.T) {
	d := GetWindow().Document()
	bodyNode := d.Body().Node()
	assert.Equal(t, bodyNode.ChildrenCount(), 3)
	el := d.CreateElement("test")
	assert.Equal(t, bodyNode.ChildrenCount(), 3)
	bodyNode.AppendChild(el.Node())
	assert.Equal(t, bodyNode.ChildrenCount(), 4)
	bodyNode.RemoveChild(el.Node())
	assert.Equal(t, bodyNode.ChildrenCount(), 3)
}
