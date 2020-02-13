package web

// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode
type AudioNode struct {
	Value
}

// GETTERS

// Context returns the associated AudioContext,
// that is the object representing the processing graph the node is participating in.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/context
func (context *AudioContext) Context() AudioContext {
	return context.Get("context").AudioContext()
}

// Inputs returns the number of inputs feeding the node.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/numberOfInputs
func (context *AudioContext) Inputs() int {
	return context.Get("numberOfInputs").Int()
}

// Outputs returns the number of outputs coming out of the node.
// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/numberOfOutputs
func (context *AudioContext) Outputs() int {
	return context.Get("numberOfOutputs").Int()
}

func (context *AudioContext) Channels() int {
	return context.Get("channelCount").Int()
}

func (context *AudioContext) ChannelsMode() ChannelsMode {
	return ChannelsMode(context.Get("channelCountMode").String())
}

func (context *AudioContext) ChannelsInterpretation() ChannelsMode {
	return ChannelsMode(context.Get("channelCountMode").String())
}

// METHODS

func (context *AudioContext) Connect(destination AudioNode, inputIndex int, outputIndex int) {
	context.Call("connect", destination.Value.Value, outputIndex, inputIndex)
}

func (context *AudioContext) DisconnectAll() {
	context.Call("disconnect")
}

func (context *AudioContext) Disconnect(destination AudioNode) {
	context.Call("disconnect", destination.Value.Value)
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
