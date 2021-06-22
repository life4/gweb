package web

import (
	"strings"
	"sync"
	"time"
)

// Object used to send HTTP requests.
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/XMLHttpRequest
type HTTPRequest struct {
	Value
}

// Send the HTTP request. This operation is blocking on the Go side
// but doesn't block JS-side main thread.
// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/send
func (req HTTPRequest) Send(body []byte) HTTPResponse {
	wg := sync.WaitGroup{}
	wg.Add(1)
	// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequestEventTarget/onload
	req.EventTarget().Listen(EventTypeLoad, func(e Event) {
		wg.Done()
	})

	if body == nil {
		req.Call("send", "")
	} else {
		// https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/atob
	}

	wg.Wait()
	return HTTPResponse{value: req.Value}
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/abort
func (req HTTPRequest) Abort() {
	req.Call("abort")
}

// https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/timeout
func (req HTTPRequest) Timeout() time.Duration {
	return time.Duration(req.Get("timeout").Int()) * time.Millisecond
}

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
	value Value
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
	return Headers(resp)
}

// Encapsulates methods to work with HTTP response headers.
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

// https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/btoa
