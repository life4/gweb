package glowasm

import (
	"syscall/js"
)

type Element struct {
	Value
}

// SUBTYPES GETTERS

func (el *Element) Attribute(namespace, name string) Attribute {
	return Attribute{value: el.Value, namespace: namespace, name: name}
}

func (el *Element) Client() Client {
	return Client{value: el.Value}
}

func (el *Element) Scroll() Scroll {
	return Scroll{value: el.Value}
}

// GETTERS

func (el *Element) AssignedSlot() string {
	return el.Get("assignedSlot").OptionalString()
}

func (el *Element) Class() string {
	return el.Get("className").String()
}

func (el *Element) Classes() []string {
	v := el.Get("classList")
	return v.Strings()
}

func (el *Element) ID() string {
	return el.Get("id").String()
}

func (el *Element) InnerHTML() string {
	return el.Get("innerHTML").String()
}

func (el *Element) LocalName() string {
	return el.Get("localName").String()
}

func (el *Element) OuterHTML() string {
	return el.Get("outerHTML").String()
}

func (el *Element) Slot() string {
	v := el.Get("slot")
	return v.OptionalString()
}

func (el *Element) TagName() string {
	return el.Get("tagName").String()
}

// POINTER METHODS

func (el *Element) ReleasePointerCapture(pointerID string) {
	el.Call("releasePointerCapture", pointerID)
}

func (el *Element) RequestPointerLock() {
	el.Call("requestPointerLock")
}

func (el *Element) SetPointerCapture(pointerID string) {
	el.Call("setPointerCapture", pointerID)
}

// OTHER METHODS

func (el *Element) Matches(selector string) bool {
	return el.Call("matches", selector).Bool()
}

func (el *Element) ScrollBy(x, y int, smooth bool) {
	if !smooth {
		el.Call("scrollBy", x, y)
		return
	}

	opts := js.Global().Get("Object").New()
	opts.Set("left", x)
	opts.Set("top", y)
	opts.Set("behavior", "smooth")
	el.Call("scrollBy", opts)
}

func (el *Element) ScrollTo(x, y int, smooth bool) {
	if !smooth {
		el.Call("scrollTo", x, y)
		return
	}

	opts := js.Global().Get("Object").New()
	opts.Set("left", x)
	opts.Set("top", y)
	opts.Set("behavior", "smooth")
	el.Call("scrollTo", opts)
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
	el.Call("scrollIntoView", opts)
}

// ELEMENT SUBTYPES

type Attribute struct {
	value     Value
	namespace string
	name      string
}

func (attr *Attribute) Value() string {
	var v Value
	if attr.namespace == "" {
		v = attr.value.Call("getAttribute", attr.name)
	} else {
		v = attr.value.Call("getAttributeNS", attr.namespace, attr.name)
	}
	return v.OptionalString()
}

func (attr *Attribute) Exists() bool {
	var v Value
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

type Client struct {
	value Value
}

func (client *Client) Height() int {
	return client.value.Get("clientHeight").Int()
}

func (client *Client) Left() int {
	return client.value.Get("clientLeft").Int()
}

func (client *Client) Top() int {
	return client.value.Get("clientTop").Int()
}

func (client *Client) Width() int {
	return client.value.Get("clientWidth").Int()
}

type Scroll struct {
	value Value
}

func (scroll *Scroll) Height() int {
	return scroll.value.Get("scrollHeight").Int()
}

func (scroll *Scroll) Left() int {
	return scroll.value.Get("scrollLeft").Int()
}

func (scroll *Scroll) Top() int {
	return scroll.value.Get("scrollTop").Int()
}

func (scroll *Scroll) Width() int {
	return scroll.value.Get("scrollWidth").Int()
}
