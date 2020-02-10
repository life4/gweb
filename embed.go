package gweb

// Embed represents HTMLEmbedElement and HTMLObjectElement.
type Embed struct {
	HTMLElement
}

func (embed *Embed) Height() int {
	return embed.Get("height").Int()
}

func (embed *Embed) Src() int {
	return embed.Get("src").Int()
}

func (embed *Embed) MIMEType() string {
	return embed.Get("type").String()
}

func (embed *Embed) Width() int {
	return embed.Get("Width").Int()
}
