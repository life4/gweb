package audio

// GainNode represents a change in volume.
// It is an AudioNode audio-processing module that causes
// a given gain to be applied to the input data before its propagation
// to the output. A GainNode always has exactly one input and one output,
// both with the same number of channels.
// https://developer.mozilla.org/en-US/docs/Web/API/GainNode
type GainNode struct {
	AudioNode
}

func (node GainNode) Gain() AudioParam {
	return AudioParam{value: node.Get("gain")}
}
