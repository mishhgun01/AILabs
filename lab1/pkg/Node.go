package pkg

type Node struct {
	st          state
	depth       int
	parent_node *Node
}

func NewNode(st state, depth int, parent *Node) *Node {
	return &Node{st: st, depth: depth, parent_node: parent}
}
