package canvas

import "syscall/js"

type Rectangle struct {
	value js.Value

	x, y, width, height int

	cleared bool
	filled  bool
	stroked bool
	corners int
}

// set params

func (rect Rectangle) Cleared() Rectangle {
	rect.cleared = true
	return rect
}

func (rect Rectangle) Filled() Rectangle {
	rect.filled = true
	return rect
}

func (rect Rectangle) Stroked() Rectangle {
	rect.stroked = true
	return rect
}

func (rect Rectangle) Rounded(radius int) Rectangle {
	rect.corners = radius
	return rect
}

// draw

func (rect Rectangle) Draw() {
	if rect.corners > 0 {
		rect.drawRoundedStroke()
		return
	}

	// clear
	if rect.cleared {
		rect.value.Call("clearRect", rect.x, rect.y, rect.width, rect.height)
		return
	}
	// stroke and fill
	if rect.stroked && rect.filled {
		rect.value.Call("rect", rect.x, rect.y, rect.width, rect.height)
		return
	}
	// only fill
	if rect.filled {
		rect.value.Call("fillRect", rect.x, rect.y, rect.width, rect.height)
		return
	}
	// only stroke (default)
	rect.value.Call("strokeRect", rect.x, rect.y, rect.width, rect.height)
}

func (rect Rectangle) drawRoundedStroke() {
	top := rect.y + rect.height
	right := rect.x + rect.width
	rad := rect.corners

	rect.value.Call("beginPath")
	rect.value.Call("moveTo", rect.x, rect.y+rad)
	rect.value.Call("lineTo", rect.x, top-rad)
	rect.value.Call("arcTo", rect.x, top, rect.x+rad, top, rad)
	rect.value.Call("lineTo", right-rad, top)
	rect.value.Call("arcTo", right, top, right, top-rad, rad)
	rect.value.Call("lineTo", right, rect.y+rad)
	rect.value.Call("arcTo", right, rect.y, right-rad, rect.y, rad)
	rect.value.Call("lineTo", rect.x+rad, rect.y)
	rect.value.Call("arcTo", rect.x, rect.y, rect.x, rect.y+rad, rad)
	rect.value.Call("stroke")
}
