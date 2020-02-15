package canvas

import "syscall/js"

type Context struct {
	js.Value
}

func (context Context) Context2D() Context2D {
	return Context2D{value: context.Value}
}
