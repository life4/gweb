package audio

// https://developer.mozilla.org/en-US/docs/Web/API/MediaStreamAudioSourceNode
type MediaStreamSourceNode struct {
	AudioNode
}

func (node MediaStreamSourceNode) Stream() MediaStream {
	return MediaStream{Value: node.Get("mediaStream")}
}
