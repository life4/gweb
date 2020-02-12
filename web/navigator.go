package gweb

type Navigator struct {
	Value
}

// PROPERTIES

func (nav Navigator) CookieEnabled() bool {
	return nav.Get("cookieEnabled").Bool()
}

func (nav Navigator) Language() string {
	return nav.Get("language").String()
}

func (nav Navigator) Languages() []string {
	return nav.Get("languages").Strings()
}

func (nav Navigator) MaxTouchPoints() int {
	return nav.Get("maxTouchPoints").Int()
}

func (nav Navigator) Online() bool {
	return nav.Get("onLine").Bool()
}

func (nav Navigator) UserAgent() string {
	return nav.Get("userAgent").String()
}

// METHODS

// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/vibrate
func (nav Navigator) Vibrate(pattern []int) {
	nav.Call("vibrate", pattern)
}
