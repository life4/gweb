package glowasm

import (
	"syscall/js"
)

type Element struct {
	value js.Value
}

// PROPERTIES

func (el *Element) AssignedSlot() string {
	v := el.value.Get("assignedSlot")
	return Value{Value: v}.OptionalString()
}

func (el *Element) Attribute(namespace, name string) Attribute {
	return Attribute{value: el.value, namespace: namespace, name: name}
}

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

func (el *Element) Slot() string {
	v := el.value.Get("slot")
	return Value{Value: v}.OptionalString()
}

func (el *Element) Tag() string {
	return el.value.Get("tagName").String()
}

// POINTER METHODS

func (el *Element) ReleasePointerCapture(pointerID string) {
	el.value.Call("releasePointerCapture", pointerID)
}

func (el *Element) RequestPointerLock() {
	el.value.Call("requestPointerLock")
}

func (el *Element) SetPointerCapture(pointerID string) {
	el.value.Call("setPointerCapture", pointerID)
}

// OTHER METHODS

func (el *Element) Matches(selector string) bool {
	return el.value.Call("matches", selector).Bool()
}

func (el *Element) ScrollBy(x, y int, smooth bool) {
	if !smooth {
		el.value.Call("scrollBy", x, y)
		return
	}

	opts := js.Global().Get("Object").New()
	opts.Set("left", x)
	opts.Set("top", y)
	opts.Set("behavior", "smooth")
	el.value.Call("scrollBy", opts)
}

func (el *Element) ScrollTo(x, y int, smooth bool) {
	if !smooth {
		el.value.Call("scrollTo", x, y)
		return
	}

	opts := js.Global().Get("Object").New()
	opts.Set("left", x)
	opts.Set("top", y)
	opts.Set("behavior", "smooth")
	el.value.Call("scrollTo", opts)
}

func (el *Element) ScrollIntoView(smooth bool, block, inline string) {
	opts := js.Global().Get("Object").New()
	opts.Set("block", block)
	opts.Set("inline", inline)
	if smooth {
		opts.Set("behavior", "smooth")
	} else {
		opts.Set("behavior", "auto")
	}
	el.value.Call("scrollIntoView", opts)
}

// ELEMENT SUBTYPES

type Attribute struct {
	value     js.Value
	namespace string
	name      string
}

func (attr *Attribute) Value() string {
	var v js.Value
	if attr.namespace == "" {
		v = attr.value.Call("getAttribute", attr.name)
	} else {
		v = attr.value.Call("getAttributeNS", attr.namespace, attr.name)
	}
	return Value{Value: v}.OptionalString()
}

func (attr *Attribute) Exists() bool {
	var v js.Value
	if attr.namespace == "" {
		v = attr.value.Call("hasAttribute", attr.name)
	} else {
		v = attr.value.Call("hasAttributeNS", attr.namespace, attr.name)
	}
	return v.Bool()
}

func (attr *Attribute) Remove() {
	if attr.namespace == "" {
		attr.value.Call("removeAttribute", attr.name)
	} else {
		attr.value.Call("removeAttributeNS", attr.namespace, attr.name)
	}
}

func (attr *Attribute) Set(value string) {
	if attr.namespace == "" {
		attr.value.Call("setAttribute", attr.name, value)
	} else {
		attr.value.Call("setAttributeNS", attr.namespace, attr.name, value)
	}
}

func (attr *Attribute) Toggle() {
	attr.value.Call("toggleAttribute", attr.name)
}
