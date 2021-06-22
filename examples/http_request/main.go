package main

import "github.com/life4/gweb/web"

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Making HTTP requests")

	// make request
	req := window.HTTPRequest("GET", "https://httpbin.org/get")
	resp := req.Send(nil)

	header := doc.CreateElement("pre")
	header.SetText(string(resp.Body()))
	body := doc.Body()
	body.Node().AppendChild(header.Node())
}
