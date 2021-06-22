package audio

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode
type AnalyserNode struct {
	AudioNode
}

func (analyser AnalyserNode) FrequencyData() FrequencyDataBytes {
	size := analyser.FFTSize()
	return FrequencyDataBytes{
		node:      analyser.Value,
		container: js.Global().Get("Uint8Array").New(size),
		Size:      size,
		Data:      make([]byte, size),
	}
}

func (analyser AnalyserNode) TimeDomain() TimeDomainBytes {
	size := analyser.FFTSize()
	return TimeDomainBytes{
		node:      analyser.Value,
		container: js.Global().Get("Uint8Array").New(size),
		Size:      size,
		Data:      make([]byte, size),
	}
}

// FFTSize represents the window size in samples that is used
// when performing a Fast Fourier Transform (FFT) to get frequency domain data.
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

// SETTERS

// FFTSize represents the window size in samples that is used
// when performing a Fast Fourier Transform (FFT) to get frequency domain data.
// https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/fftSize
func (analyser AnalyserNode) SetFFTSize(value int) {
	analyser.Set("fftSize", value)
}

// https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/minDecibels
func (analyser AnalyserNode) SetMinDecibels(value int) {
	analyser.Set("minDecibels", value)
}

// https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/maxDecibels
func (analyser AnalyserNode) SetMaxDecibels(value int) {
	analyser.Set("maxDecibels", value)
}

// https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/smoothingTimeConstant
func (analyser AnalyserNode) SetSmoothingTimeConstant(value float64) {
	analyser.Set("smoothingTimeConstant", value)
}

// SUBTYPES

type TimeDomainBytes struct {
	node      js.Value // AnalyserNode value to do method call
	container js.Value // where to read data in JS
	Size      int      // Size of the data array
	Data      []byte   // where to copy data from JS into Go
}

// Update reads the current waveform or time-domain into `Data` attribute.
func (domain *TimeDomainBytes) Update() {
	domain.node.Call("getByteTimeDomainData", domain.container)
	js.CopyBytesToGo(domain.Data, domain.container)
}

type FrequencyDataBytes struct {
	node      js.Value // AnalyserNode value to do method call
	container js.Value // where to read data in JS
	Size      int      // Size of the data array
	Data      []byte   // where to copy data from JS into Go
}

// Update reads the current frequency data into `Data` attribute.
// The frequency data is composed of integers on a scale from 0 to 255.
func (freq *FrequencyDataBytes) Update() {
	freq.node.Call("getByteFrequencyData", freq.container)
	js.CopyBytesToGo(freq.Data, freq.container)
}
