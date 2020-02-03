package glowasm

import (
	"syscall/js"
)

type Element struct {
	value js.Value
}

func (el *Element) Class() string {
	return el.value.Get("className").String()
}

func (el *Element) Classes() []string {
	value := Value{Value: el.value.Get("classList")}
	return value.Strings()
}
