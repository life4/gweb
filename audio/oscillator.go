package audio

// https://developer.mozilla.org/en-US/docs/Web/API/OscillatorNode
type OscillatorNode struct {
	AudioNode
}

// PROPERTIES

func (node OscillatorNode) Frequency() AudioParam {
	return AudioParam{value: node.Get("frequency")}
}

func (node OscillatorNode) DeTune() AudioParam {
	return AudioParam{value: node.Get("detune")}
}

// Shape specifies the shape of waveform to play.
// https://developer.mozilla.org/en-US/docs/Web/API/OscillatorNode/type
func (node OscillatorNode) Shape() Shape {
	return Shape(node.Get("type").String())
}

func (node OscillatorNode) SetShape(shape Shape) {
	node.Set("type", string(shape))
}

// METHODS

// https://developer.mozilla.org/en-US/docs/Web/API/OscillatorNode/start
func (node OscillatorNode) Start(when float64) {
	node.Call("start", when)
}

// https://developer.mozilla.org/en-US/docs/Web/API/OscillatorNode/stop
func (node OscillatorNode) Stop(when float64) {
	node.Call("stop", when)
}

// SUBTYPES

type Shape string

const (
	ShapeSine     = Shape("sine")
	ShapeSquare   = Shape("square")
	ShapeSawTooth = Shape("sawtooth")
	ShapeTriangle = Shape("triangle")
	ShapeCustom   = Shape("custom")
)
