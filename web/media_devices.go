package web

type MediaDevices struct {
	Value
}

// https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices/getUserMedia
func (devices MediaDevices) Audio() MediaStream {
	params := map[string]interface{}{"audio": true}
	return devices.Call("getUserMedia", params).MediaStream()
}

// https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices/getUserMedia
func (devices MediaDevices) Video() MediaStream {
	params := map[string]interface{}{"video": true}
	return devices.Call("getUserMedia", params).MediaStream()
}
