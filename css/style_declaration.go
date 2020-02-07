package css

import "syscall/js"

type CSSStyleDeclaration struct {
	Value js.Value
}

// RULES MANIPULATION

func (decl *CSSStyleDeclaration) Len() int {
	return decl.Value.Get("length").Int()
}

func (decl *CSSStyleDeclaration) Names() []string {
	length := decl.Len()
	items := make([]string, length, 0)
	for i := 0; i < length; i++ {
		items[i] = decl.Value.Call("item", i).String()
	}
	return items
}

func (decl *CSSStyleDeclaration) Get(name string) string {
	return decl.Value.Call("getPropertyValue", name).String()
}

func (decl *CSSStyleDeclaration) Set(name, value string, important bool) {
	priority := ""
	if important {
		priority = "important"
	}
	decl.Value.Call("setProperty", name, value, priority)
}

func (decl *CSSStyleDeclaration) Remove(name string) {
	decl.Value.Call("removeProperty", name)
}

func (decl *CSSStyleDeclaration) Important(name string) bool {
	return decl.Value.Call("getPropertyPriority", name).String() == "important"
}

// RULES GETTERS

func (decl *CSSStyleDeclaration) Background() string {
	return decl.Get("background")
}

func (decl *CSSStyleDeclaration) BackgroundAttachment() string {
	return decl.Get("background-attachment")
}

func (decl *CSSStyleDeclaration) BackgroundColor() string {
	return decl.Get("background-color")
}

func (decl *CSSStyleDeclaration) BackgroundImage() string {
	return decl.Get("background-image")
}

func (decl *CSSStyleDeclaration) BackgroundPosition() string {
	return decl.Get("background-position")
}

func (decl *CSSStyleDeclaration) BackgroundRepeat() string {
	return decl.Get("background-repeat")
}

func (decl *CSSStyleDeclaration) Border() string {
	return decl.Get("border")
}

func (decl *CSSStyleDeclaration) BorderBottom() string {
	return decl.Get("border-bottom")
}

func (decl *CSSStyleDeclaration) BorderBottomColor() string {
	return decl.Get("border-bottom-color")
}

func (decl *CSSStyleDeclaration) BorderBottomStyle() string {
	return decl.Get("border-bottom-style")
}

func (decl *CSSStyleDeclaration) BorderBottomWidth() string {
	return decl.Get("border-bottom-width")
}

func (decl *CSSStyleDeclaration) BorderColor() string {
	return decl.Get("border-color")
}

func (decl *CSSStyleDeclaration) BorderLeft() string {
	return decl.Get("border-left")
}

func (decl *CSSStyleDeclaration) BorderLeftColor() string {
	return decl.Get("border-left-color")
}

func (decl *CSSStyleDeclaration) BorderLeftStyle() string {
	return decl.Get("border-left-style")
}

func (decl *CSSStyleDeclaration) BorderLeftWidth() string {
	return decl.Get("border-left-width")
}

func (decl *CSSStyleDeclaration) BorderRight() string {
	return decl.Get("border-right")
}

func (decl *CSSStyleDeclaration) BorderRightColor() string {
	return decl.Get("border-right-color")
}

func (decl *CSSStyleDeclaration) BorderRightStyle() string {
	return decl.Get("border-right-style")
}

func (decl *CSSStyleDeclaration) BorderRightWidth() string {
	return decl.Get("border-right-width")
}

func (decl *CSSStyleDeclaration) BorderStyle() string {
	return decl.Get("border-style")
}

func (decl *CSSStyleDeclaration) BorderTop() string {
	return decl.Get("border-top")
}

func (decl *CSSStyleDeclaration) BorderTopColor() string {
	return decl.Get("border-top-color")
}

func (decl *CSSStyleDeclaration) BorderTopStyle() string {
	return decl.Get("border-top-style")
}

func (decl *CSSStyleDeclaration) BorderTopWidth() string {
	return decl.Get("border-top-width")
}

func (decl *CSSStyleDeclaration) BorderWidth() string {
	return decl.Get("border-width")
}

func (decl *CSSStyleDeclaration) Clear() string {
	return decl.Get("clear")
}

func (decl *CSSStyleDeclaration) Clip() string {
	return decl.Get("clip")
}

func (decl *CSSStyleDeclaration) Color() string {
	return decl.Get("color")
}

func (decl *CSSStyleDeclaration) Cursor() string {
	return decl.Get("cursor")
}

func (decl *CSSStyleDeclaration) Display() string {
	return decl.Get("display")
}

func (decl *CSSStyleDeclaration) Filter() string {
	return decl.Get("filter")
}

func (decl *CSSStyleDeclaration) CssFloat() string {
	return decl.Get("float")
}

func (decl *CSSStyleDeclaration) Font() string {
	return decl.Get("font")
}

func (decl *CSSStyleDeclaration) FontFamily() string {
	return decl.Get("font-family")
}

func (decl *CSSStyleDeclaration) FontSize() string {
	return decl.Get("font-size")
}

func (decl *CSSStyleDeclaration) FontVariant() string {
	return decl.Get("font-variant")
}

func (decl *CSSStyleDeclaration) FontWeight() string {
	return decl.Get("font-weight")
}

func (decl *CSSStyleDeclaration) Height() string {
	return decl.Get("height")
}

func (decl *CSSStyleDeclaration) Left() string {
	return decl.Get("left")
}

func (decl *CSSStyleDeclaration) LetterSpacing() string {
	return decl.Get("letter-spacing")
}

func (decl *CSSStyleDeclaration) LineHeight() string {
	return decl.Get("line-height")
}

func (decl *CSSStyleDeclaration) ListStyle() string {
	return decl.Get("list-style")
}

func (decl *CSSStyleDeclaration) ListStyleImage() string {
	return decl.Get("list-style-image")
}

func (decl *CSSStyleDeclaration) ListStylePosition() string {
	return decl.Get("list-style-position")
}

func (decl *CSSStyleDeclaration) ListStyleType() string {
	return decl.Get("list-style-type")
}

func (decl *CSSStyleDeclaration) Margin() string {
	return decl.Get("margin")
}

func (decl *CSSStyleDeclaration) MarginBottom() string {
	return decl.Get("margin-bottom")
}

func (decl *CSSStyleDeclaration) MarginLeft() string {
	return decl.Get("margin-left")
}

func (decl *CSSStyleDeclaration) MarginRight() string {
	return decl.Get("margin-right")
}

func (decl *CSSStyleDeclaration) MarginTop() string {
	return decl.Get("margin-top")
}

func (decl *CSSStyleDeclaration) Overflow() string {
	return decl.Get("overflow")
}

func (decl *CSSStyleDeclaration) Padding() string {
	return decl.Get("padding")
}

func (decl *CSSStyleDeclaration) PaddingBottom() string {
	return decl.Get("padding-bottom")
}

func (decl *CSSStyleDeclaration) PaddingLeft() string {
	return decl.Get("padding-left")
}

func (decl *CSSStyleDeclaration) PaddingRight() string {
	return decl.Get("padding-right")
}

func (decl *CSSStyleDeclaration) PaddingTop() string {
	return decl.Get("padding-top")
}

func (decl *CSSStyleDeclaration) PageBreakAfter() string {
	return decl.Get("page-break-after")
}

func (decl *CSSStyleDeclaration) PageBreakBefore() string {
	return decl.Get("page-break-before")
}

func (decl *CSSStyleDeclaration) Position() string {
	return decl.Get("position")
}

func (decl *CSSStyleDeclaration) StrokeDasharray() string {
	return decl.Get("stroke-dasharray")
}

func (decl *CSSStyleDeclaration) StrokeDashoffset() string {
	return decl.Get("stroke-dashoffset")
}

func (decl *CSSStyleDeclaration) StrokeWidth() string {
	return decl.Get("stroke-width")
}

func (decl *CSSStyleDeclaration) TextAlign() string {
	return decl.Get("text-align")
}

func (decl *CSSStyleDeclaration) TextDecoration() string {
	return decl.Get("text-decoration")
}

func (decl *CSSStyleDeclaration) TextIndent() string {
	return decl.Get("text-indent")
}

func (decl *CSSStyleDeclaration) TextTransform() string {
	return decl.Get("text-transform")
}

func (decl *CSSStyleDeclaration) Top() string {
	return decl.Get("top")
}

func (decl *CSSStyleDeclaration) VerticalAlign() string {
	return decl.Get("vertical-align")
}

func (decl *CSSStyleDeclaration) Visibility() string {
	return decl.Get("visibility")
}

func (decl *CSSStyleDeclaration) Width() string {
	return decl.Get("width")
}

func (decl *CSSStyleDeclaration) ZIndex() string {
	return decl.Get("z-index")
}

// SETTERS

func (decl *CSSStyleDeclaration) SetBackground(value string, important bool) {
	decl.Set("background", value, important)
}

func (decl *CSSStyleDeclaration) SetBackgroundAttachment(value string, important bool) {
	decl.Set("background-attachment", value, important)
}

func (decl *CSSStyleDeclaration) SetBackgroundColor(value string, important bool) {
	decl.Set("background-color", value, important)
}

func (decl *CSSStyleDeclaration) SetBackgroundImage(value string, important bool) {
	decl.Set("background-image", value, important)
}

func (decl *CSSStyleDeclaration) SetBackgroundPosition(value string, important bool) {
	decl.Set("background-position", value, important)
}

func (decl *CSSStyleDeclaration) SetBackgroundRepeat(value string, important bool) {
	decl.Set("background-repeat", value, important)
}

func (decl *CSSStyleDeclaration) SetBorder(value string, important bool) {
	decl.Set("border", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderBottom(value string, important bool) {
	decl.Set("border-bottom", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderBottomColor(value string, important bool) {
	decl.Set("border-bottom-color", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderBottomStyle(value string, important bool) {
	decl.Set("border-bottom-style", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderBottomWidth(value string, important bool) {
	decl.Set("border-bottom-width", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderColor(value string, important bool) {
	decl.Set("border-color", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderLeft(value string, important bool) {
	decl.Set("border-left", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderLeftColor(value string, important bool) {
	decl.Set("border-left-color", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderLeftStyle(value string, important bool) {
	decl.Set("border-left-style", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderLeftWidth(value string, important bool) {
	decl.Set("border-left-width", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderRight(value string, important bool) {
	decl.Set("border-right", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderRightColor(value string, important bool) {
	decl.Set("border-right-color", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderRightStyle(value string, important bool) {
	decl.Set("border-right-style", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderRightWidth(value string, important bool) {
	decl.Set("border-right-width", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderStyle(value string, important bool) {
	decl.Set("border-style", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderTop(value string, important bool) {
	decl.Set("border-top", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderTopColor(value string, important bool) {
	decl.Set("border-top-color", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderTopStyle(value string, important bool) {
	decl.Set("border-top-style", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderTopWidth(value string, important bool) {
	decl.Set("border-top-width", value, important)
}

func (decl *CSSStyleDeclaration) SetBorderWidth(value string, important bool) {
	decl.Set("border-width", value, important)
}

func (decl *CSSStyleDeclaration) SetClear(value string, important bool) {
	decl.Set("clear", value, important)
}

func (decl *CSSStyleDeclaration) SetClip(value string, important bool) {
	decl.Set("clip", value, important)
}

func (decl *CSSStyleDeclaration) SetColor(value string, important bool) {
	decl.Set("color", value, important)
}

func (decl *CSSStyleDeclaration) SetCursor(value string, important bool) {
	decl.Set("cursor", value, important)
}

func (decl *CSSStyleDeclaration) SetDisplay(value string, important bool) {
	decl.Set("display", value, important)
}

func (decl *CSSStyleDeclaration) SetFilter(value string, important bool) {
	decl.Set("filter", value, important)
}

func (decl *CSSStyleDeclaration) SetCssFloat(value string, important bool) {
	decl.Set("float", value, important)
}

func (decl *CSSStyleDeclaration) SetFont(value string, important bool) {
	decl.Set("font", value, important)
}

func (decl *CSSStyleDeclaration) SetFontFamily(value string, important bool) {
	decl.Set("font-family", value, important)
}

func (decl *CSSStyleDeclaration) SetFontSize(value string, important bool) {
	decl.Set("font-size", value, important)
}

func (decl *CSSStyleDeclaration) SetFontVariant(value string, important bool) {
	decl.Set("font-variant", value, important)
}

func (decl *CSSStyleDeclaration) SetFontWeight(value string, important bool) {
	decl.Set("font-weight", value, important)
}

func (decl *CSSStyleDeclaration) SetHeight(value string, important bool) {
	decl.Set("height", value, important)
}

func (decl *CSSStyleDeclaration) SetLeft(value string, important bool) {
	decl.Set("left", value, important)
}

func (decl *CSSStyleDeclaration) SetLetterSpacing(value string, important bool) {
	decl.Set("letter-spacing", value, important)
}

func (decl *CSSStyleDeclaration) SetLineHeight(value string, important bool) {
	decl.Set("line-height", value, important)
}

func (decl *CSSStyleDeclaration) SetListStyle(value string, important bool) {
	decl.Set("list-style", value, important)
}

func (decl *CSSStyleDeclaration) SetListStyleImage(value string, important bool) {
	decl.Set("list-style-image", value, important)
}

func (decl *CSSStyleDeclaration) SetListStylePosition(value string, important bool) {
	decl.Set("list-style-position", value, important)
}

func (decl *CSSStyleDeclaration) SetListStyleType(value string, important bool) {
	decl.Set("list-style-type", value, important)
}

func (decl *CSSStyleDeclaration) SetMargin(value string, important bool) {
	decl.Set("margin", value, important)
}

func (decl *CSSStyleDeclaration) SetMarginBottom(value string, important bool) {
	decl.Set("margin-bottom", value, important)
}

func (decl *CSSStyleDeclaration) SetMarginLeft(value string, important bool) {
	decl.Set("margin-left", value, important)
}

func (decl *CSSStyleDeclaration) SetMarginRight(value string, important bool) {
	decl.Set("margin-right", value, important)
}

func (decl *CSSStyleDeclaration) SetMarginTop(value string, important bool) {
	decl.Set("margin-top", value, important)
}

func (decl *CSSStyleDeclaration) SetOverflow(value string, important bool) {
	decl.Set("overflow", value, important)
}

func (decl *CSSStyleDeclaration) SetPadding(value string, important bool) {
	decl.Set("padding", value, important)
}

func (decl *CSSStyleDeclaration) SetPaddingBottom(value string, important bool) {
	decl.Set("padding-bottom", value, important)
}

func (decl *CSSStyleDeclaration) SetPaddingLeft(value string, important bool) {
	decl.Set("padding-left", value, important)
}

func (decl *CSSStyleDeclaration) SetPaddingRight(value string, important bool) {
	decl.Set("padding-right", value, important)
}

func (decl *CSSStyleDeclaration) SetPaddingTop(value string, important bool) {
	decl.Set("padding-top", value, important)
}

func (decl *CSSStyleDeclaration) SetPageBreakAfter(value string, important bool) {
	decl.Set("page-break-after", value, important)
}

func (decl *CSSStyleDeclaration) SetPageBreakBefore(value string, important bool) {
	decl.Set("page-break-before", value, important)
}

func (decl *CSSStyleDeclaration) SetPosition(value string, important bool) {
	decl.Set("position", value, important)
}

func (decl *CSSStyleDeclaration) SetStrokeDasharray(value string, important bool) {
	decl.Set("stroke-dasharray", value, important)
}

func (decl *CSSStyleDeclaration) SetStrokeDashoffset(value string, important bool) {
	decl.Set("stroke-dashoffset", value, important)
}

func (decl *CSSStyleDeclaration) SetStrokeWidth(value string, important bool) {
	decl.Set("stroke-width", value, important)
}

func (decl *CSSStyleDeclaration) SetTextAlign(value string, important bool) {
	decl.Set("text-align", value, important)
}

func (decl *CSSStyleDeclaration) SetTextDecoration(value string, important bool) {
	decl.Set("text-decoration", value, important)
}

func (decl *CSSStyleDeclaration) SetTextIndent(value string, important bool) {
	decl.Set("text-indent", value, important)
}

func (decl *CSSStyleDeclaration) SetTextTransform(value string, important bool) {
	decl.Set("text-transform", value, important)
}

func (decl *CSSStyleDeclaration) SetTop(value string, important bool) {
	decl.Set("top", value, important)
}

func (decl *CSSStyleDeclaration) SetVerticalAlign(value string, important bool) {
	decl.Set("vertical-align", value, important)
}

func (decl *CSSStyleDeclaration) SetVisibility(value string, important bool) {
	decl.Set("visibility", value, important)
}

func (decl *CSSStyleDeclaration) SetWidth(value string, important bool) {
	decl.Set("width", value, important)
}

func (decl *CSSStyleDeclaration) SetZIndex(value string, important bool) {
	decl.Set("z-index", value, important)
}
