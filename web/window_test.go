package web

import (
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWindow(t *testing.T) {
	w := GetWindow()
	assert.Equal(t, w.Type(), js.TypeObject, "window is undefined")
	assert.Equal(t, w.Call("toString").String(), "[object Window]", "bad type")
	assert.Equal(t, w.Get("asdqwqwafd").Type(), js.TypeUndefined, "your types lie")
}

func TestWindowConsole(t *testing.T) {
	c := GetWindow().Console()
	assert.Equal(t, c.Type(), js.TypeObject, "console is undefined")
}

func TestWindowDocument(t *testing.T) {
	d := GetWindow().Document()
	assert.Equal(t, d.Type(), js.TypeObject, "document is undefined")
	assert.Equal(t, d.Call("toString").String(), "[object HTMLDocument]", "bad type")
}
