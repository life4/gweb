package web

import (
	"github.com/life4/gweb/css"
)

type Editable string

const (
	EditableTrue    = Editable("true")
	EditableFalse   = Editable("false")
	EditableInherit = Editable("inherit")
)

type Direction string

const (
	DirectionLTR  = Direction("ltr")
	DirectionRTL  = Direction("rtl")
	DirectionAuto = Direction("auto")
)

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement
type HTMLElement struct {
	Element
}

// SUBTYPES GETTERS

// Incapsulates a set of offset-related properties.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetHeight
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetLeft
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetParent
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetTop
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetWidth
func (el *HTMLElement) Offset() Offset {
	return Offset{value: el.Value}
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/style
func (el *HTMLElement) Style() css.CSSStyleDeclaration {
	value := el.Get("style").JSValue()
	return css.CSSStyleDeclaration{Value: value}
}

// GETTERS

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dir
func (el *HTMLElement) Direction() Direction {
	return Direction(el.Get("dir").String())
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/isContentEditable
func (el *HTMLElement) Editable() bool {
	return el.Get("isContentEditable").Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/hidden
func (el *HTMLElement) Hidden() bool {
	return el.Get("hidden").Bool()
}

func (el *HTMLElement) Lang() string {
	return el.Get("lang").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLOrForeignElement/nonce
func (el *HTMLElement) Nonce() string {
	return el.Get("nonce").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
func (el *HTMLElement) Text() string {
	return el.Get("innerText").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLOrForeignElement/tabIndex
func (el *HTMLElement) TabIndex() int {
	return el.Get("tabIndex").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/title
func (el *HTMLElement) Title() string {
	return el.Get("title").String()
}

// SETTERS

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dir
func (el HTMLElement) SetDirection(value Direction) {
	el.Set("dir", string(value))
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/contentEditable
func (el HTMLElement) SetEditable(value Editable) {
	el.Set("contentEditable", string(value))
}

func (el HTMLElement) SetHidden(value bool) {
	el.Set("hidden", value)
}

func (el HTMLElement) SetLang(value string) {
	el.Set("lang", value)
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
func (el HTMLElement) SetText(text string) {
	el.Set("innerText", text)
}

// HTMLElement SUBTYPES

// Incapsulates a set of offset-related properties.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetHeight
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetLeft
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetParent
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetTop
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetWidth
type Offset struct {
	value Value
}

// Returns the height of an element, relative to the layout.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetHeight
func (offset *Offset) Height() int {
	return offset.value.Get("offsetHeight").Int()
}

// Returns the distance from this element's left border to its Offset.Parent's left border.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetLeft
func (offset *Offset) Left() int {
	return offset.value.Get("offsetLeft").Int()
}

// Returns the distance from this element's top border to its Offset.Parent's top border.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetTop
func (offset *Offset) Top() int {
	return offset.value.Get("offsetTop").Int()
}

// Returns the width of an element, relative to the layout.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetWidth
func (offset *Offset) Width() int {
	return offset.value.Get("offsetWidth").Int()
}

// Returns an Element that is the element
// from which all offset calculations are currently computed.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetParent
func (offset *Offset) Parent() Element {
	v := offset.value.Get("offsetParent")
	return Element{Value: v}
}
