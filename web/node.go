package web

import "syscall/js"

type NodeType int

const (
	ELEMENT_NODE                = NodeType(1)
	ATTRIBUTE_NODE              = NodeType(2) // deprecated
	TEXT_NODE                   = NodeType(3)
	CDATA_SECTION_NODE          = NodeType(4)
	ENTITY_REFERENCE_NODE       = NodeType(5) // deprecated
	ENTITY_NODE                 = NodeType(6) // deprecated
	PROCESSING_INSTRUCTION_NODE = NodeType(7)
	COMMENT_NODE                = NodeType(8)
	DOCUMENT_NODE               = NodeType(9)
	DOCUMENT_TYPE_NODE          = NodeType(10)
	DOCUMENT_FRAGMENT_NODE      = NodeType(11)
	NOTATION_NODE               = NodeType(12) // deprecated
)

type Node struct {
	value Value
}

// PROPERTIES

// Returns the base URL of the document containing the Node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/baseURI
func (node Node) BaseURI() string {
	return node.value.Get("baseURI").String()
}

func (node Node) ChildrenCount() int {
	return node.value.Get("childElementCount").Int()
}

func (node Node) Connected() bool {
	return node.value.Get("isConnected").Bool()
}

func (node Node) Content() string {
	return node.value.Get("textContent").String()
}

func (node Node) Document() Document {
	value := node.value.Get("ownerDocument")
	switch value.Type() {
	case js.TypeNull:
		return Document{Value: node.value}
	default:
		return Document{Value: value}
	}
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeName
func (node Node) Name() string {
	return node.value.Get("nodeName").String()
}

// Returns the type of the node.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeType
func (node Node) Type() NodeType {
	return NodeType(node.value.Get("nodeType").Int())
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeValue
func (node Node) Value() string {
	return node.value.Get("nodeValue").OptionalString()
}

// METHODS

// Clean up all the text nodes under this element (merge adjacent, remove empty).
// https://developer.mozilla.org/en-US/docs/Web/API/Node/normalize
func (node Node) Normalize() {
	node.value.Call("normalize")
}

func (node Node) Clone(deep bool) Node {
	return node.value.Call("cloneNode", deep).Node()
}

// TREE

// https://developer.mozilla.org/en-US/docs/Web/API/Node/appendChild
func (node Node) AppendChild(child Node) {
	node.value.Call("appendChild", child.value)
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/childNodes
func (node Node) ChildNodes() []HTMLElement {
	nodes := node.value.Get("childNodes")
	values := nodes.Values()
	elements := make([]HTMLElement, len(values))
	for i, value := range values {
		elements[i] = value.HTMLElement()
	}
	return elements
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/firstChild
func (node Node) FirstChild() HTMLElement {
	return node.value.Get("firstChild").HTMLElement()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/hasChildNodes
func (node Node) HasChildNodes() bool {
	return node.value.Call("hasChildNodes").Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/parentElement
func (node Node) Parent() HTMLElement {
	return node.value.Get("parentElement").HTMLElement()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/removeChild
func (node Node) RemoveChild(child Node) {
	node.value.Call("removeChild", child.value)
}

// Remove all children.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/removeChild
func (node Node) RemoveChildren() {
	for {
		child := node.FirstChild()
		if child.Type() == js.TypeNull {
			return
		}
		node.value.Call("removeChild", child.Value)
	}
}

// Remove the node from the parent node.
func (node Node) Remove() bool {
	parent := node.Parent()
	if parent.Type() == js.TypeNull {
		return false
	}
	parent.Call("removeChild", node.value)
	return true
}
