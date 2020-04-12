package web

// Embed represents HTMLEmbedElement and HTMLObjectElement.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLEmbedElement
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement
type Embed struct {
	HTMLElement
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/height
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object#attr-height
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed#attr-height
func (embed *Embed) Height() int {
	return embed.Get("height").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed#attr-src
func (embed *Embed) Src() int {
	return embed.Get("src").Int()
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/type
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object#attr-type
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed#attr-type
func (embed *Embed) MIMEType() string {
	return embed.Get("type").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLObjectElement/width
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object#attr-width
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed#attr-width
func (embed *Embed) Width() int {
	return embed.Get("Width").Int()
}
