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

// ATTRIBUTE METHODS

func (el *Element) GetAttribute(namespace, name string) string {
	var v js.Value
	if namespace == "" {
		v = el.value.Call("getAttribute", name)
	} else {
		v = el.value.Call("getAttributeNS", namespace, name)
	}
	return Value{Value: v}.OptionalString()
}

func (el *Element) HasAttribute(namespace, name string) bool {
	var v js.Value
	if namespace == "" {
		v = el.value.Call("hasAttribute", name)
	} else {
		v = el.value.Call("hasAttributeNS", namespace, name)
	}
	return v.Bool()
}

func (el *Element) RemoveAttribute(namespace, name string) bool {
	var v js.Value
	if namespace == "" {
		v = el.value.Call("removeAttribute", name)
	} else {
		v = el.value.Call("removeAttributeNS", namespace, name)
	}
	return v.Bool()
}

func (el *Element) SetAttribute(namespace, name, value string) bool {
	var v js.Value
	if namespace == "" {
		v = el.value.Call("setAttribute", name, value)
	} else {
		v = el.value.Call("setAttributeNS", namespace, name, value)
	}
	return v.Bool()
}

// POINTER METHODS

func (el *Element) ReleasePointerCapture(pointerID string) {
	el.value.Call("releasePointerCapture", pointerID)
}

func (el *Element) RequestPointerLock() {
	el.value.Call("requestPointerLock")
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
