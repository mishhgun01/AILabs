package node

import "AILabs/lab1/pkg/state"

type Node struct {
	State state.State
	Depth int

	pathCost      int
	costGenerator func(node *Node)
	parent        *Node
}

func NewNode(depth int, state state.State, f pathCostGenerator, parent *Node) *Node {
	node := &Node{State: state, Depth: depth, costGenerator: f, parent: parent}
	f(node)
	return node
}

func (node *Node) GenerateSubnodes() []Node {
	generatedStates := node.State.GenerateSubstates()
	var newNodes []Node
	for _, st := range generatedStates {
		newNode := *NewNode(node.Depth+1, st, node.costGenerator, node)
		newNodes = append(newNodes, newNode)
	}
	return newNodes
}

func (node *Node) ChainSlice() []Node {
	var chain []Node
	temp := node
	for temp != nil {
		chain = append([]Node{*temp}, chain...)
		temp = temp.parent
	}
	return chain
}
