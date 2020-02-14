package audio

// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode
type AudioNode struct {
	Value
}

// GETTERS

// Context returns the associated AudioContext,
// that is the object representing the processing graph the node is participating in.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/context
func (node AudioNode) Context() AudioContext {
	return AudioContext{Value: node.Get("context")}
}

// Inputs returns the number of inputs feeding the node.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/numberOfInputs
func (node AudioNode) Inputs() int {
	return node.Get("numberOfInputs").Int()
}

// Outputs returns the number of outputs coming out of the node.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/numberOfOutputs
func (node AudioNode) Outputs() int {
	return node.Get("numberOfOutputs").Int()
}

func (node AudioNode) Channels() int {
	return node.Get("channelCount").Int()
}

func (node AudioNode) ChannelsMode() ChannelsMode {
	return ChannelsMode(node.Get("channelCountMode").String())
}

func (node AudioNode) ChannelsInterpretation() ChannelsMode {
	return ChannelsMode(node.Get("channelCountMode").String())
}

// METHODS

func (node AudioNode) Connect(destination AudioNode, inputIndex int, outputIndex int) {
	node.Call("connect", destination.Value.Value, outputIndex, inputIndex)
}

func (node AudioNode) DisconnectAll() {
	node.Call("disconnect")
}

func (node AudioNode) Disconnect(destination AudioNode) {
	node.Call("disconnect", destination.Value.Value)
}

// SUBTYPES

type Channels struct {
	value Value
}

// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/channelCount
func (channels Channels) Count() int {
	return channels.value.Get("channelCount").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/channelCountMode
func (channels Channels) Mode() ChannelsMode {
	return ChannelsMode(channels.value.Get("channelCountMode").String())
}

// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/channelInterpretation
func (channels Channels) Discrete() bool {
	return channels.value.Get("channelInterpretation").String() == "discrete"
}

// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/channelInterpretation
func (channels Channels) Speakers() bool {
	return channels.value.Get("channelInterpretation").String() == "speakers"
}

type ChannelsMode string

const (
	ChannelsModeMax        = ChannelsMode("max")
	ChannelsModeClampedMax = ChannelsMode("clamped-max")
	ChannelsModeExplicit   = ChannelsMode("explicit")
)
