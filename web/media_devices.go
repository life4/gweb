package web

type MediaDevices struct {
	Value
}

// https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices/getUserMedia
func (devices MediaDevices) Audio() Promise {
	params := map[string]any{"audio": true}
	return devices.Call("getUserMedia", params).Promise()
}

// https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices/getUserMedia
func (devices MediaDevices) Video() Promise {
	params := map[string]any{"video": true}
	return devices.Call("getUserMedia", params).Promise()
}
