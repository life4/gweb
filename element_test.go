package glowasm

import (
	"syscall/js"
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

func TestElementClass(t *testing.T) {
	element := GetWindow().Document().CreateElement("", "lol")
	class := element.Class()

	assert.Equal(t, class.String(), "")
	class.Set("one two")
	assert.Equal(t, class.String(), "one two")

	assert.Equal(t, class.Strings(), []string{"one", "two"})
	class.Append("three", "four")
	assert.Equal(t, class.Strings(), []string{"one", "two", "three", "four"})
	class.Remove("two", "three")
	assert.Equal(t, class.Strings(), []string{"one", "four"})
	assert.Equal(t, class.String(), "one four")

	assert.True(t, class.Contains("one"))
	assert.False(t, class.Contains("two"))
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

func TestElementSlots(t *testing.T) {
	d := GetWindow().Document()
	body := d.Body()

	// create <template>
	template := d.CreateElement("", "template")
	body.Node().AppendChild(template.Node())

	// add <slot> into template
	slot := d.CreateElement("", "slot")
	slot.Set("name", "example")
	slot.SetInnerHTML("default text")
	template.Node().AppendChild(slot.Node())

	// make <span> element that will fill the <slot>
	span := d.CreateElement("", "span")
	assert.Equal(t, span.Slot(), "")
	span.SetSlot("example")
	assert.Equal(t, span.Slot(), "example")
	body.Node().AppendChild(span.Node())

	// render template
	shadow := body.Shadow().Attach()
	assert.Equal(t, span.AssignedSlot().Type(), js.TypeNull)
	shadow.Node().AppendChild(template.Content())
	assert.NotEqual(t, span.AssignedSlot().Type(), js.TypeNull)
	assert.Equal(t, span.AssignedSlot().InnerHTML(), "default text")

	// clean up
	assert.Equal(t, body.Node().ChildrenCount(), 4)
	span.Node().Remove()
	assert.Equal(t, body.Node().ChildrenCount(), 3)
}
