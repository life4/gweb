package audio

import "syscall/js"

type MediaStream struct {
	js.Value
}

// Casts audio.MediaStream to js.Value
func (stream MediaStream) JSValue() js.Value {
	return stream.Value
}

// PROPERTIES

// https://developer.mozilla.org/en-US/docs/Web/API/MediaStream/active
func (stream MediaStream) Active() bool {
	return stream.Get("active").Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/MediaStream/id
func (stream MediaStream) ID() string {
	return stream.Get("active").String()
}

// METHODS

// https://developer.mozilla.org/en-US/docs/Web/API/MediaStream/clone
func (stream MediaStream) Clone() MediaStream {
	return MediaStream{Value: stream.Call("clone")}
}
