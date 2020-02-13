package audio

// https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext
// https://developer.mozilla.org/en-US/docs/Web/API/AudioContext
type AudioContext struct {
	Value
}

// GETTERS

// Current time returns an ever-increasing hardware time in seconds used for scheduling. It starts at 0.
// https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/currentTime
func (context AudioContext) CurrentTime() float64 {
	return context.Get("currentTime").Float()
}

// SampleRate returns the sample rate (in samples per second) used by all nodes in this context.
// https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/sampleRate
func (context AudioContext) SampleRate() int {
	return context.Get("sampleRate").Int()
}

func (context AudioContext) State() AudioContextState {
	return AudioContextState(context.Get("state").String())
}

// METHODS

func (context AudioContext) Analyser() AnalyserNode {
	return context.Call("createAnalyser").AnalyserNode()
}

// SUBTYPES

type AudioContextState string

const (
	AudioContextStateSuspended = AudioContextState("suspended")
	AudioContextStateRunning   = AudioContextState("running")
	AudioContextStateClosed    = AudioContextState("closed")
)
