package glowasm

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

func (node *Node) ChildNodes() []Element {
	nodes := node.value.Get("childNodes")
	values := nodes.Values()
	elements := make([]Element, len(values))
	for i, value := range values {
		elements[i] = Element{Value: value}
	}
	return elements
}

func (node *Node) HasChildNodes() bool {
	return node.value.Call("hasChildNodes").Bool()
}

func (node *Node) Parent() Element {
	value := node.value.Get("parentElement")
	switch value.Type() {
	case js.TypeNull:
		return Element{}
	default:
		return Element{Value: value}
	}
}