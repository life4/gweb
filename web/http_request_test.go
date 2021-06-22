package web

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPRequest_GET(t *testing.T) {
	is := require.New(t)
	req := GetWindow().HTTPRequest("GET", "https://httpbin.org/get")
	resp := req.Send(nil)
	is.Equal(resp.StatusCode(), 200)
	is.Equal(resp.Status(), "")
	is.Equal(resp.Headers().Get("Content-Type"), "application/json")
	is.Equal(resp.Headers().Values(), []string{"content-length: 619", "content-type: application/json"})
}
