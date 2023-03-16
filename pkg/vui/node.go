package vui

import (
	"strings"
)

// Interface port from DOM WebAPI
// https://developer.mozilla.org/en-US/docs/Web/API/Node
// Note: The line comment out means not implement yet
type INode interface {
	AppendChild(...INode) error
	CloneNode(bool) INode
	Contains(INode) bool
	GetRootNode() INode
	HasChildNodes() bool
	InsertBefore()
	IsDefaultNamespace() bool
	IsEqualNode() bool
	IsSameNode() bool
	Normalize()
	RemoveChild()
	ReplaceChild()
	// CompareDocumentPosition()
	// LookupNamespaceURI()
	// LookupPrefix()

	/* Getters for properties */
	ChildNodes() NodeList
	IsConnected() bool
	NextSibling() INode
	PreviousSibling() INode
	OwnerDocument() Documentary
	NodeName() string
	NodeValue() string

	/* Setter for properties */
	SetOwnerDocument(Documentary) error
	SetTextContent(string)
	bindTo(INode) error
}

type NodeList []INode

var (
	_ INode = (*Node)(nil)
)

func (l NodeList) First() INode {
	if len(l) > 0 {
		return l[0]
	}
	return nil
}

func (l NodeList) Last() INode {
	if len(l) > 0 {
		return l[len(l)-1]
	}
	return nil
}

type NodeOpt func(*Node)

// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeName
func WithNodeName(name string, nodeType NodeType) NodeOpt {
	var nodeName string
	switch nodeType {
	case AttributeNode:
		nodeName = "#attr" // @TODO
	case CDATASectionNode:
		nodeName = "#cdata-section"
	case CommentNode:
		nodeName = "#comment"
	case DocumentNode:
		nodeName = "#document"
	case DocumentFragmentNode:
		nodeName = "#document-fragment"
	case DocumentTypeNode:
		nodeName = "#documentType.name" // @TODO
	case ElementNode:
		nodeName = strings.ToUpper(name)
	case ProcessingInstructionNode:
		nodeName = "#proccessionInstruction.target" // @TODO
	case TextNode:
		nodeName = "#text"
	}

	return NodeOpt(func(n *Node) {
		n.nodeName = nodeName
		n.nodeType = nodeType
	})
}

func WithNodeValue(val string) NodeOpt {
	return NodeOpt(func(n *Node) {
		n.nodeValue = val
	})
}

func NewNode(opts ...NodeOpt) *Node {
	n := &Node{
		childNodes: make(NodeList, 0),
	}

	for _, opt := range opts {
		opt(n)
	}

	return n
}

// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeType
type NodeType uint32

const (
	ErrorNode                 NodeType = 0
	ElementNode                        = 1
	AttributeNode                      = 2
	TextNode                           = 3
	CDATASectionNode                   = 4
	ProcessingInstructionNode          = 7
	CommentNode                        = 8
	DocumentNode                       = 9
	DocumentTypeNode                   = 10
	DocumentFragmentNode               = 11
)

type Node struct {
	childNodes NodeList
	nodeName   string
	nodeValue  string
	nodeType   NodeType

	ownerDocument   Documentary
	textContent     string
	nextSibling     INode
	parentNode      INode
	previousSibling INode
}

func (n *Node) SetTextContent(content string) {
	n.textContent = content
}

func (n *Node) IsConnected() bool {
	return n.parentNode != nil
}

func (n *Node) NextSibling() INode {
	return n.nextSibling
}

func (n *Node) PreviousSibling() INode {
	return n.previousSibling
}

func (n *Node) OwnerDocument() Documentary {
	return n.ownerDocument
}

func (n *Node) SetOwnerDocument(doc Documentary) error {
	n.ownerDocument = doc
	return nil
}

func (n *Node) bindTo(parent INode) error {
	n.parentNode = parent
	// n.ownerDocument = parent.OwnerDocument()
	return nil
}

// ChildNodes return the Nodelist of children's node
func (n *Node) ChildNodes() NodeList {
	return n.childNodes
}

// ChildNodes return the Nodelist of children's node
func (n *Node) NodeName() string {
	return n.nodeName
}

// ChildNodes return the Nodelist of children's node
func (n *Node) NodeValue() string {
	return n.nodeValue
}

// AppendChild() adds a node to the end of the children node.
// If the given child is a reference to an existing node in the document,
// appendChild() moves it from its current position to the new position.
func (n *Node) AppendChild(nl ...INode) error {
	for _, child := range nl {
		// Resolve parent, previousSibling, nextSibling
		if err := child.bindTo(n); err != nil {
			return err
		}
		n.childNodes = append(n.childNodes, child)
	}
	return nil
}

func (n *Node) CloneNode(deep bool) INode {
	new := *n
	new.parentNode = nil
	return &new
}

func (n *Node) Contains(node INode) bool {
	// @TODO
	return true
}

func (n *Node) GetRootNode() INode {
	// The GetRootNode() method of the Node interface returns the context object's root,
	// which optionally includes the shadow root if it is available.
	// @TODO
	return n
}

func (n *Node) HasChildNodes() bool {
	// @TODO
	return true
}

func (n *Node) InsertBefore() {
	// @TODO
}

func (n *Node) IsDefaultNamespace() bool {
	// @TODO
	return true
}

func (n *Node) IsEqualNode() bool {
	// @TODO
	return true
}

func (n *Node) IsSameNode() bool {
	// @TODO
	return true
}

func (n *Node) Normalize() {
	// @TODO
}

func (n *Node) RemoveChild() {
	// @TODO
}

func (n *Node) ReplaceChild() {
	// @TODO
}
