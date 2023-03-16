package vui

type INode interface {
	appendChild(INode)
	cloneNode(INode)
	contains()
	getRootNode()
	hasChildNodes()
	insertBefore()
	isDefaultNamespace()
	isEqualNode()
	isSameNode()
	normalize()
	removeChild()
	replaceChild()
}

type NodeList []INode

func (l *NodeList) first() *INode {
	if len(l) > 0 {
		return l[0]
	}
	return nil
}

func (l *NodeList) last() *INode {
	if len(l) > 0 {
		return l[-1]
	}
	return nil
}

type Node struct {
	childNodes       NodeList
	firstChild       INode
	isConnected      bool
	lastChild        INode
	nextSibling      INode
	nodeName         string
	nodeValue        interface{}
	parentNode       INode
	previousSibiling INode
	textContent      string
}

func (n *Node) appendChild(nd INode) {
	n.childNodes = append(n.childNodes, nd)
}

func (n *Node) cloneNode() {
	new := Node{
		childNo
	}

}

func (n *Node) contains() {}

func (n *Node) getRootNode() {}

func (n *Node) hasChildNodes() {}

func (n *Node) insertBefore() {}

func (n *Node) isDefaultNamespace() {}

func (n *Node) isEqualNode() {}

func (n *Node) isSameNode() {}

func (n *Node) normalize() {}

func (n *Node) removeChild() {}

func (n *Node) replaceChild() {}
