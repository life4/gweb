package web

// https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext
// https://developer.mozilla.org/en-US/docs/Web/API/AudioContext
type AudioContext struct {
	Value
}

// getters

func (context *AudioContext) CurrentTime() float64 {
	return context.Get("currentTime").Float()
}

func (context *AudioContext) SampleRate() int {
	return context.Get("sampleRate").Int()
}

func (context *AudioContext) State() AudioContextState {
	return AudioContextState(context.Get("state").String())
}

// SUBTYPES

type AudioContextState string

const (
	AudioContextStateSuspended = AudioContextState("suspended")
	AudioContextStateRunning   = AudioContextState("running")
	AudioContextStateClosed    = AudioContextState("closed")
)
