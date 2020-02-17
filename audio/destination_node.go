package audio

type DestinationNode struct {
	AudioNode
}

func (node DestinationNode) MaxChannels() int {
	return node.Get("maxChannelCount").Int()
}
