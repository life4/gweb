package glowasm

type Editable string

const (
	EDITABLE_TRUE    = Editable("true")
	EDITABLE_FALSE   = Editable("false")
	EDITABLE_INHERIT = Editable("inherit")
)

type Direction string

const (
	DIRECTION_LTR  = Direction("ltr")
	DIRECTION_RTL  = Direction("rtl")
	DIRECTION_AUTO = Direction("auto")
)

type HTMLElement struct {
	Element
}

// SUBTYPES GETTERS

func (el *HTMLElement) Node() Node {
	return Node{value: el.Value}
}

func (el *HTMLElement) Offset() Offset {
	return Offset{value: el.Value}
}

// GETTERS

func (el *HTMLElement) Direction() Direction {
	return Direction(el.Get("dir").String())
}

func (el *HTMLElement) Editable() bool {
	return el.Get("isContentEditable").Bool()
}

func (el *HTMLElement) Hidden() bool {
	return el.Get("hidden").Bool()
}

func (el *HTMLElement) Lang() string {
	return el.Get("lang").String()
}

func (el *HTMLElement) Nonce() string {
	return el.Get("nonce").String()
}

func (el *HTMLElement) Text() string {
	return el.Get("innerText").String()
}

func (el *HTMLElement) TabIndex() int {
	return el.Get("tabIndex").Int()
}

func (el *HTMLElement) Title() string {
	return el.Get("title").String()
}

// SETTERS

func (el *HTMLElement) SetDirection(value Editable) {
	el.Set("dir", string(value))
}

func (el *HTMLElement) SetEditable(value Editable) {
	el.Set("contentEditable", string(value))
}

func (el *HTMLElement) SetHidden(value bool) {
	el.Set("hidden", value)
}

func (el *HTMLElement) SetLang(value string) {
	el.Set("lang", value)
}

// HTMLElement SUBTYPES

type Offset struct {
	value Value
}

func (offset *Offset) Height() int {
	return offset.value.Get("offsetHeight").Int()
}

func (offset *Offset) Left() int {
	return offset.value.Get("offsetLeft").Int()
}

func (offset *Offset) Top() int {
	return offset.value.Get("offsetTop").Int()
}

func (offset *Offset) Width() int {
	return offset.value.Get("offsetWidth").Int()
}

func (offset *Offset) Parent() Element {
	v := offset.value.Get("offsetParent")
	return Element{Value: v}
}
