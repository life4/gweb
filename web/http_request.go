package web

import (
	"strings"
	"sync"
	"syscall/js"
	"time"
)

// Object used to send HTTP requests.
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/XMLHttpRequest
type HTTPRequest struct {
	Value
	window Window
}

// Send the HTTP request. This operation is blocking on the Go side
// but doesn't block JS-side main thread.
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/send
func (req HTTPRequest) Send(body []byte) HTTPResponse {
	wg := sync.WaitGroup{}
	wg.Add(1)
	// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/onreadystatechange
	req.Set("onreadystatechange", js.FuncOf(func(this js.Value, args []js.Value) any {
		state := req.Get("readyState").Int()
		if state == 4 || state == 0 {
			wg.Done()
		}
		return nil
	}))

	if body == nil {
		req.Call("send", nil)
	} else {
		encoded := req.window.Get("Uint8Array").New(len(body))
		js.CopyBytesToJS(encoded.Value, body)
		req.Call("send", encoded)
	}

	wg.Wait()
	return HTTPResponse{
		value:  req.Value,
		window: req.window,
	}
}

// Abort aborts the request if it has already been sent.
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/abort
func (req HTTPRequest) Abort() {
	req.Call("abort")
}

// Timeout represents how long a request can take before automatically being terminated.
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/timeout
func (req HTTPRequest) Timeout() time.Duration {
	return time.Duration(req.Get("timeout").Int()) * time.Millisecond
}

// SetTimeout sets the time after which the request will be terminated.
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/timeout
func (req HTTPRequest) SetTimeout(timeout time.Duration) {
	req.Set("timeout", int(timeout/time.Millisecond))
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/withCredentials
func (req HTTPRequest) WithCredentials() bool {
	return req.Get("withCredentials").Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/withCredentials
func (req HTTPRequest) SetWithCredentials(creds bool) {
	req.Set("withCredentials", creds)
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/setRequestHeader
func (req HTTPRequest) SetHeader(header, value string) {
	req.Call("setRequestHeader", header, value)
}

type HTTPResponse struct {
	value  Value
	window Window
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/response
func (resp HTTPResponse) Body() []byte {
	raw := resp.value.Get("response")
	if raw.IsNull() {
		return nil
	}
	raw = resp.window.Get("Uint8Array").New(raw)
	dec := make([]byte, raw.Length())
	js.CopyBytesToGo(dec, raw.Value)
	return dec
}

// Finished indicates is the request is succesfully completed.
// It can be false if the request was aborted.
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/readyState
func (resp HTTPResponse) Finished() bool {
	return resp.value.Get("readyState").Int() == 4
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/responseText
func (resp HTTPResponse) Text() string {
	return resp.value.Get("responseText").OptionalString()
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/responseURL
func (resp HTTPResponse) URL() string {
	return resp.value.Get("responseURL").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/status
func (resp HTTPResponse) StatusCode() int {
	return resp.value.Get("status").Int()
}

// Always an empty string for HTTP/2 responses.
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/statusText
func (resp HTTPResponse) Status() string {
	return resp.value.Get("statusText").String()
}

func (resp HTTPResponse) Headers() Headers {
	return Headers{value: resp.value}
}

// Headers encapsulates methods to work with HTTP response headers.
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/getAllResponseHeaders
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/getResponseHeader
type Headers struct {
	value Value
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/getResponseHeader
func (h Headers) Get(name string) string {
	return h.value.Call("getResponseHeader", name).OptionalString()
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/getResponseHeader
func (h Headers) Has(name string) bool {
	return !h.value.Call("getResponseHeader", name).IsNull()
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/getAllResponseHeaders
func (h Headers) Values() []string {
	vals := h.value.Call("getAllResponseHeaders").String()
	return strings.Split(strings.TrimSpace(vals), "\r\n")
}
