package node

import "AILabs/lab1/pkg/state"

type Node struct {
	State       state.State
	Depth       int
	path_cost   int
	parent_node *Node
}

func NewNode(depth int, state state.State, f PathCostFunction, parent_Node *Node) *Node {
	return &Node{State: state, Depth: depth, path_cost: f(parent_Node), parent_node: parent_Node}

}

func (node *Node) GenerateSubNodes(f PathCostFunction) []Node {
	generated_states := node.State.GenerateSubstates()
	var newNodes []Node
	for _, st := range generated_states {
		newNodes = append(newNodes, *NewNode(node.Depth+1, st, f, node))
	}
	return newNodes
}

func (node *Node) GetChain() []Node {
	var chain []Node
	temp := node
	for temp != nil {
		chain = append([]Node{*temp}, chain...)
		temp = temp.parent_node
	}
	return chain
}
