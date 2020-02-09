package glowasm

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentURL(t *testing.T) {
	d := GetWindow().Document()
	assert.True(t, strings.HasPrefix(d.URL(), "http://localhost:"), "bad URL")
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
	assert.Equal(t, d.Domain(), "localhost")
}

func TestDocumentReferrer(t *testing.T) {
	d := GetWindow().Document()
	assert.Equal(t, d.Referrer(), "")
}

func TestDocumentInputEncoding(t *testing.T) {
	d := GetWindow().Document()
	assert.Equal(t, d.InputEncoding(), "UTF-8")
}

func TestDocumentReadyState(t *testing.T) {
	d := GetWindow().Document()
	assert.Equal(t, d.ReadyState(), "complete")
}

func TestDocumentTitle(t *testing.T) {
	d := GetWindow().Document()
	assert.Equal(t, d.Title(), "Go wasm")
}
