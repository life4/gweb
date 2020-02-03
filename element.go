package glowasm

import (
	"syscall/js"
)

type Element struct {
	value js.Value
}

// PROPERTIES

func (el *Element) Class() string {
	return el.value.Get("className").String()
}

func (el *Element) Classes() []string {
	value := Value{Value: el.value.Get("classList")}
	return value.Strings()
}

func (el *Element) ClientHeight() int {
	return el.value.Get("clientHeight").Int()
}

func (el *Element) ClientLeft() int {
	return el.value.Get("clientLeft").Int()
}

func (el *Element) ClientTop() int {
	return el.value.Get("clientTop").Int()
}

func (el *Element) ClientWidth() int {
	return el.value.Get("ClientWidth").Int()
}

func (el *Element) ID() string {
	return el.value.Get("id").String()
}

func (el *Element) InnerHTML() string {
	return el.value.Get("innerHTML").String()
}

func (el *Element) LocalName() string {
	return el.value.Get("localName").String()
}

func (el *Element) OuterHTML() string {
	return el.value.Get("outerHTML").String()
}

func (el *Element) ScrollHeight() int {
	return el.value.Get("scrollHeight").Int()
}

func (el *Element) ScrollLeft() int {
	return el.value.Get("scrollLeft").Int()
}

func (el *Element) ScrollTop() int {
	return el.value.Get("scrollTop").Int()
}

func (el *Element) ScrollWidth() int {
	return el.value.Get("scrollWidth").Int()
}

func (el *Element) Tag() string {
	return el.value.Get("tagName").String()
}

// METHODS

func (el *Element) Attribute(name string) string {
	v := el.value.Get("getAttribute")
	return Value{Value: v}.OptionalString()
}
