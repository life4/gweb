package glowasm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElementAttribute(t *testing.T) {
	b := GetWindow().Document().Body()
	attr := b.Attribute("", "test")
	assert.False(t, attr.Exists())

	attr.Set("val")
	assert.True(t, attr.Exists())
	assert.Equal(t, attr.Get(), "val")

	attr.Remove()
	assert.False(t, attr.Exists())
}

func TestElementClient(t *testing.T) {
	b := GetWindow().Document().Body()
	c := b.Client()
	assert.Equal(t, c.Width(), 784)
	assert.Equal(t, c.Height(), 0)
	assert.Equal(t, c.Left(), 0)
	assert.Equal(t, c.Top(), 0)
}

func TestElementScroll(t *testing.T) {
	b := GetWindow().Document().Body()
	s := b.Scroll()
	assert.Equal(t, s.Width(), 784)
	assert.Equal(t, s.Height(), 0)
	assert.Equal(t, s.Left(), 0)
	assert.Equal(t, s.Top(), 0)
}
