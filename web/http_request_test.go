package web

import (
	"encoding/json"
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

func TestHTTPRequest_POST(t *testing.T) {
	is := require.New(t)
	req := GetWindow().HTTPRequest("POST", "https://httpbin.org/post")
	resp := req.Send([]byte("hello world"))
	is.Equal(resp.StatusCode(), 200)
	is.Equal(resp.Status(), "")
	is.Equal(resp.Headers().Get("Content-Type"), "application/json")
	is.Equal(resp.Headers().Values(), []string{"content-length: 772", "content-type: application/json"})

	data := struct {
		Data string `json:"data"`
		URL  string `json:"url"`
	}{}
	err := json.Unmarshal(resp.Body(), &data)
	is.Nil(err)
	is.Equal(data.URL, "https://httpbin.org/post")
	is.Equal(data.Data, "hello world")
}
