package web

import (
	"syscall/js"

	"github.com/life4/gweb/audio"
)

// Value is an extended js.Value with more types support
type Value struct {
	js.Value
}

// overloaded methods

// Call calls the given method of the object.
func (v Value) Call(method string, args ...any) Value {
	result := v.Value.Call(method, unwrapValues(args)...)
	return Value{Value: result}
}

// Get returns the given property of the value
func (v Value) Get(property string) Value {
	result := v.Value.Get(property)
	return Value{Value: result}
}

// Creates new instance of the JS class.
func (v Value) New(args ...any) Value {
	result := v.Value.New(unwrapValues(args)...)
	return Value{Value: result}
}

// new methods

// Casts web.Value to js.Value
func (v Value) JSValue() js.Value {
	return v.Value
}

// Represents the current value into Canvas
func (v Value) Canvas() Canvas {
	return Canvas{HTMLElement: v.HTMLElement()}
}

// Represents the current value into Element
func (v Value) Element() Element {
	return Element{Value: v}
}

// Represents the current value into Embed
func (v Value) Embed() Embed {
	return Embed{HTMLElement: v.HTMLElement()}
}

// Represents the current value into Event
func (v Value) Event() Event {
	return Event{Value: v}
}

// Represents the current value into EventTarget
func (v Value) EventTarget() EventTarget {
	return EventTarget{Value: v}
}

// Represents the current value into HTMLElement
func (v Value) HTMLElement() HTMLElement {
	return HTMLElement{Element: v.Element()}
}

// Represents the current value into audio.MediaStream
func (v Value) MediaStream() audio.MediaStream {
	return audio.MediaStream{Value: v.Value}
}

// Represents the current value into Node
func (v Value) Node() Node {
	return Node{value: v}
}

// Represents the current value into Promise
func (v Value) Promise() Promise {
	return Promise{Value: v}
}

// Represents the current value as slice of values
func (v *Value) Values() (items []Value) {
	len := v.Get("length").Int()
	for i := 0; i < len; i++ {
		item := v.Call("item", i)
		items = append(items, item)
	}
	return items
}

// Represents the current value as slice of strings
func (v Value) Strings() (items []string) {
	len := v.Get("length").Int()
	for i := 0; i < len; i++ {
		item := v.Call("item", i)
		items = append(items, item.String())
	}
	return items
}

// OptionalString returns empty string if Value is null
func (v Value) OptionalString() string {
	switch v.Type() {
	case js.TypeNull:
		return ""
	case js.TypeString:
		return v.String()
	default:
		panic("bad type")
	}
}

func unwrapValues(args []any) []any {
	values := make([]any, len(args))
	for i, arg := range args {
		switch arg := arg.(type) {
		case Value:
			values[i] = arg.Value
		default:
			values[i] = arg
		}
	}
	return values
}
