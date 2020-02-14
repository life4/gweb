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

func (node *Node) BaseURI() string {
	return node.value.Get("baseURI").String()
}

func (node Node) ChildrenCount() int {
	return node.value.Get("childElementCount").Int()
}

func (node *Node) Connected() bool {
	return node.value.Get("isConnected").Bool()
}

func (node *Node) Content() string {
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

func (node *Node) Name() string {
	return node.value.Get("nodeName").String()
}

func (node *Node) Type() NodeType {
	return NodeType(node.value.Get("nodeType").Int())
}

func (node *Node) Value() string {
	return node.value.Get("nodeValue").OptionalString()
}

// METHODS

func (node *Node) Normalize() {
	node.value.Call("normalize")
}

// TREE

func (node Node) AppendChild(child Node) {
	node.value.Call("appendChild", child.value)
}

func (node Node) ChildNodes() []HTMLElement {
	nodes := node.value.Get("childNodes")
	values := nodes.Values()
	elements := make([]HTMLElement, len(values))
	for i, value := range values {
		elements[i] = value.HTMLElement()
	}
	return elements
}

func (node Node) FirstChild() HTMLElement {
	return node.value.Get("firstChild").HTMLElement()
}

func (node Node) HasChildNodes() bool {
	return node.value.Call("hasChildNodes").Bool()
}

func (node Node) Parent() HTMLElement {
	return node.value.Get("parentElement").HTMLElement()
}

func (node Node) RemoveChild(child Node) {
	node.value.Call("removeChild", child.value)
}

func (node Node) RemoveChildren() {
	for {
		child := node.FirstChild()
		if child.Type() == js.TypeNull {
			return
		}
		node.value.Call("removeChild", child.Value)
	}
}

func (node Node) Remove() bool {
	parent := node.Parent()
	if parent.Type() == js.TypeNull {
		return false
	}
	parent.Call("removeChild", node.value)
	return true
}
