package web

type Navigator struct {
	Value
}

// PROPERTIES

// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/cookieEnabled
func (nav Navigator) CookieEnabled() bool {
	return nav.Get("cookieEnabled").Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage/language
func (nav Navigator) Language() string {
	return nav.Get("language").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage/languages
func (nav Navigator) Languages() []string {
	return nav.Get("languages").Strings()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/maxTouchPoints
func (nav Navigator) MaxTouchPoints() int {
	return nav.Get("maxTouchPoints").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/mediaDevices
func (nav Navigator) MediaDevices() MediaDevices {
	return MediaDevices{Value: nav.Get("mediaDevices")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/onLine
func (nav Navigator) Online() bool {
	return nav.Get("onLine").Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/NavigatorID/userAgent
func (nav Navigator) UserAgent() string {
	return nav.Get("userAgent").String()
}

// METHODS

// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/vibrate
func (nav Navigator) Vibrate(pattern []int) {
	nav.Call("vibrate", pattern)
}
