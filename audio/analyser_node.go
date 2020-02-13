package audio

// https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode
type AnalyserNode struct {
	AudioNode
}

// FFTSize represents the window size in samples that is used
// when performing a Fast Fourier Transform (FFT) to get frequency domain data..
// https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/fftSize
func (analyser AnalyserNode) FFTSize() int {
	return analyser.Get("fftSize").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/frequencyBinCount
func (analyser AnalyserNode) FrequencyBinCount() int {
	return analyser.Get("frequencyBinCount").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/minDecibels
func (analyser AnalyserNode) MinDecibels() int {
	return analyser.Get("minDecibels").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/msxDecibels
func (analyser AnalyserNode) MaxDecibels() int {
	return analyser.Get("maxDecibels").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/smoothingTimeConstant
func (analyser AnalyserNode) SmoothingTimeConstant() float64 {
	return analyser.Get("smoothingTimeConstant").Float()
}
