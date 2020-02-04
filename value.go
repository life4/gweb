package glowasm

import (
	"syscall/js"
)

// Value is an extended js.Value with more types support
type Value struct {
	js.Value
}

// overloaded methods

func (v Value) Call(method string, args ...interface{}) Value {
	result := v.Value.Call(method, args...)
	return Value{Value: result}
}

func (v Value) Get(property string) Value {
	result := v.Value.Get(property)
	return Value{Value: result}
}

// new methods

func (v *Value) Canvas() Canvas {
	return Canvas{Element: v.Element()}
}

func (v *Value) Element() Element {
	return Element{Value: *v}
}

func (v *Value) Values() (items []Value) {
	len := v.Get("length").Int()
	for i := 0; i < len; i++ {
		item := v.Call("item", i)
		items = append(items, item)
	}
	return items
}

func (v *Value) Strings() (items []string) {
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
