package web

import (
	"syscall/js"
)

type Element struct {
	Value
}

// SUBTYPES GETTERS

func (el *Element) Attribute(name string) Attribute {
	return Attribute{value: el.Value, Namespace: "", Name: name}
}

func (el *Element) Class() Class {
	return Class{value: el.Value}
}

func (el *Element) Client() Client {
	return Client{value: el.Value}
}

func (el Element) Shadow() ShadowDOM {
	return ShadowDOM{value: el.Value}
}

func (el *Element) Scroll() Scroll {
	return Scroll{value: el.Value}
}

// SLOTS

func (el Element) AssignedSlot() Element {
	return el.Get("assignedSlot").Element()
}

func (el Element) Slot() string {
	return el.Get("slot").OptionalString()
}

func (el Element) SetSlot(name string) {
	el.Set("slot", name)
}

// GETTERS

func (el *Element) ID() string {
	return el.Get("id").String()
}

func (el Element) InnerHTML() string {
	return el.Get("innerHTML").String()
}

func (el *Element) LocalName() string {
	return el.Get("localName").String()
}

func (el *Element) OuterHTML() string {
	return el.Get("outerHTML").String()
}

func (el *Element) TagName() string {
	return el.Get("tagName").String()
}

// SETTERS

func (el Element) SetID(id string) {
	el.Set("id", id)
}

func (el Element) SetInnerHTML(html string) {
	el.Set("innerHTML", html)
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
	Namespace string
	Name      string
}

func (attr *Attribute) Get() string {
	var v Value
	if attr.Namespace == "" {
		v = attr.value.Call("getAttribute", attr.Name)
	} else {
		v = attr.value.Call("getAttributeNS", attr.Namespace, attr.Name)
	}
	return v.OptionalString()
}

func (attr Attribute) Exists() bool {
	var v Value
	if attr.Namespace == "" {
		v = attr.value.Call("hasAttribute", attr.Name)
	} else {
		v = attr.value.Call("hasAttributeNS", attr.Namespace, attr.Name)
	}
	return v.Bool()
}

func (attr Attribute) Remove() {
	if attr.Namespace == "" {
		attr.value.Call("removeAttribute", attr.Name)
	} else {
		attr.value.Call("removeAttributeNS", attr.Namespace, attr.Name)
	}
}

func (attr Attribute) Set(value string) {
	if attr.Namespace == "" {
		attr.value.Call("setAttribute", attr.Name, value)
	} else {
		attr.value.Call("setAttributeNS", attr.Namespace, attr.Name, value)
	}
}

func (attr Attribute) Toggle() {
	attr.value.Call("toggleAttribute", attr.Name)
}

type Client struct {
	value Value
}

func (client Client) Height() int {
	return client.value.Get("clientHeight").Int()
}

func (client Client) Left() int {
	return client.value.Get("clientLeft").Int()
}

func (client Client) Top() int {
	return client.value.Get("clientTop").Int()
}

func (client Client) Width() int {
	return client.value.Get("clientWidth").Int()
}

type Scroll struct {
	value Value
}

func (scroll Scroll) Height() int {
	return scroll.value.Get("scrollHeight").Int()
}

func (scroll Scroll) Left() int {
	return scroll.value.Get("scrollLeft").Int()
}

func (scroll Scroll) Top() int {
	return scroll.value.Get("scrollTop").Int()
}

func (scroll Scroll) Width() int {
	return scroll.value.Get("scrollWidth").Int()
}

type ShadowDOM struct {
	value Value
}

// Attach attaches a shadow DOM tree to the specified element and returns ShadowRoot.
// We always create "open" shadow DOM because "closed" can't totally
// forbid access to the DOM and give falls feeling of protection.
// Read more: https://blog.revillweb.com/open-vs-closed-shadow-dom-9f3d7427d1af
func (shadow ShadowDOM) Attach() Element {
	opts := js.Global().Get("Object").New()
	opts.Set("mode", "open")
	return shadow.value.Call("attachShadow", opts).Element()
}

// Host returns a reference to the DOM element the ShadowRoot is attached to.
func (shadow ShadowDOM) Host() Element {
	return shadow.value.Get("host").Element()
}

// Root returns ShadowRoot hosted by the element.
func (shadow ShadowDOM) Root() Element {
	return shadow.value.Get("shadowRoot").Element()
}

type Class struct {
	value Value
}

// String returns `class` attribute
func (cls Class) String() string {
	return cls.value.Get("className").String()
}

// Strings returns classes from `class` attribute
func (cls Class) Strings() []string {
	v := cls.value.Get("classList")
	return v.Strings()
}

// Contains returns true if `class` attribute contains given class
func (cls Class) Contains(name string) bool {
	return cls.value.Get("classList").Call("contains", name).Bool()
}

// Add adds new class into `class` attribute
func (cls Class) Append(names ...string) {
	if len(names) == 0 {
		return
	}
	casted := make([]interface{}, len(names))
	for i, name := range names {
		casted[i] = interface{}(name)
	}
	cls.value.Get("classList").Call("add", casted...)
}

// Remove removes class from classes list in `class` attribute
func (cls Class) Remove(names ...string) {
	if len(names) == 0 {
		return
	}
	casted := make([]interface{}, len(names))
	for i, name := range names {
		casted[i] = interface{}(name)
	}
	cls.value.Get("classList").Call("remove", casted...)
}

// Set overwrites the whole `class` attribute
func (cls Class) Set(name string) {
	cls.value.Set("className", name)
}
