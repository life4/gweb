package glowasm

import (
	"syscall/js"
)

type Value struct {
	value js.Value
}

func (v *Value) Canvas() Canvas {
	return Canvas{value: v.value}
}
