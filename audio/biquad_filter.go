package audio

// https://developer.mozilla.org/en-US/docs/Web/API/BiquadFilterNode
type BiquadFilterNode struct {
	AudioNode
}

// https://developer.mozilla.org/en-US/docs/Web/API/BiquadFilterNode/frequency
func (node BiquadFilterNode) Frequency() AudioParam {
	return AudioParam{value: node.Get("frequency")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/BiquadFilterNode/detune
func (node BiquadFilterNode) DeTune() AudioParam {
	return AudioParam{value: node.Get("detune")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/BiquadFilterNode/gain
func (node BiquadFilterNode) Gain() AudioParam {
	return AudioParam{value: node.Get("gain")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/BiquadFilterNode/Q
func (node BiquadFilterNode) QFactor() AudioParam {
	return AudioParam{value: node.Get("Q")}
}

// FilterType is kind of filtering algorithm the node is implementing
// https://developer.mozilla.org/en-US/docs/Web/API/BiquadFilterNode/type
func (node BiquadFilterNode) FilterType() FilterType {
	return FilterType(node.Get("type").String())
}

// SUBTYPES

type FilterType string

const (
	FilterTypeLowPass   = FilterType("lowpass")
	FilterTypeHighPass  = FilterType("highpass")
	FilterTypeBandPass  = FilterType("bandpass")
	FilterTypeLowShelf  = FilterType("lowshelf")
	FilterTypeHighShelf = FilterType("highshelf")
	FilterTypePeaking   = FilterType("peaking")
	FilterTypeNotch     = FilterType("notch")
	FilterTypeAllPass   = FilterType("allpass")
)
