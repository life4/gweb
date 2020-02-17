package audio

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/AudioParam
type AudioParam struct {
	value js.Value
}

func (param AudioParam) Default() float64 {
	return param.value.Get("defaultValue").Float()
}

func (param AudioParam) Max() float64 {
	return param.value.Get("maxValue").Float()
}

func (param AudioParam) Min() float64 {
	return param.value.Get("minValue").Float()
}

func (param AudioParam) Get() float64 {
	return param.value.Get("value").Float()
}

func (param AudioParam) Set(value float64) {
	param.value.Set("value", value)
}

// AtTime returns a namespace of operations on AudioParam
// that are scheduled at specified time.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioParam#Methods
func (param AudioParam) AtTime(time float64) AtTime {
	return AtTime{value: param.value, time: time}
}

// AtTime is a namespace of operations on AudioParam
// that are scheduled at specified time.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioParam#Methods
type AtTime struct {
	value js.Value
	time  float64
}

// Set schedules an instant change to the AudioParam value at a precise time,
// as measured against AudioContext.CurrentTime.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioParam/setValueAtTime
func (param AtTime) Set(value float64) {
	param.value.Call("setValueAtTime", value, param.time)
}

// LinearRampTo schedules a gradual linear change in the value of the AudioParam.
// The change starts at the time specified for the previous event,
// follows a linear ramp to the new value given in the value parameter,
// and reaches the new value at the time given in the `time` parameter.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioParam/linearRampToValueAtTime
func (param AtTime) LinearRampTo(value float64) {
	param.value.Call("linearRampToValueAtTime", value, param.time)
}

// ExponentialRampTo schedules a gradual exponential change in the value of the AudioParam.
// The change starts at the time specified for the previous event,
// follows an exponential ramp to the new value given in the value parameter,
// and reaches the new value at the time given in the `time` parameter.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioParam/exponentialRampToValueAtTime
func (param AtTime) ExponentialRampTo(value float64) {
	param.value.Call("exponentialRampToValueAtTime", value, param.time)
}

// SetTarget schedules the start of a gradual change to the AudioParam value.
// This is useful for decay or release portions of ADSR envelopes.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioParam/setTargetAtTime
func (param AtTime) SetTarget(target, timeConstant float64) {
	param.value.Call("setTargetAtTime", target, param.time, timeConstant)
}

// SetCurve schedules the parameter's value to change following a curve
// defined by a list of values. The curve is a linear interpolation between
// the sequence of values defined in an array of floating-point values,
// which are scaled to fit into the given interval starting at `time` and a specific `duration`.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioParam/setValueCurveAtTime
func (param AtTime) SetCurve(values []float64, duration float64) {
	param.value.Call("setValueCurveAtTime", values, param.time, duration)
}

// Cancel cancels all scheduled future changes to the AudioParam.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioParam/cancelScheduledValues
func (param AtTime) Cancel(values []float64, duration float64) {
	param.value.Call("cancelScheduledValues", param.time)
}

// CancelAndHold cancels all scheduled future changes to the AudioParam
// but holds its value at a given time until further changes are made using other methods.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioParam/cancelAndHoldAtTime
func (param AtTime) CancelAndHold() {
	param.value.Call("cancelAndHoldAtTime", param.time)
}
