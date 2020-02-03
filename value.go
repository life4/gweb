package glowasm

import (
	"syscall/js"
)

// Value is an extended js.Value with more types support
type Value struct {
	js.Value
}

func (v *Value) Canvas() Canvas {
	el := Element{Value: v.Value}
	return Canvas{Element: el}
}

func (v *Value) Values() (items []Value) {
	len := v.Get("length").Int()
	for i := 0; i < len; i++ {
		item := v.Call("item", i)
		items = append(items, Value{Value: item})
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
