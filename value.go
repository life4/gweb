package glowasm

import (
	"syscall/js"
)

type Value struct {
	js.Value
}

func (v *Value) Canvas() Canvas {
	return Canvas{value: v.Value}
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
