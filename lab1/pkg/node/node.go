package node

import "AILabs/lab1/pkg/state"

type Node struct {
	State state.State
	Depth int

	pathCost      int
	costGenerator func(node *Node) int
	parent        *Node
}

func NewNode(depth int, state state.State, f pathCostGenerator, parent *Node) *Node {
	return &Node{State: state, Depth: depth, pathCost: f(parent), costGenerator: f, parent: parent}

}

func (node *Node) GenerateSubnodes() []Node {
	generatedStates := node.State.GenerateSubstates()
	var newNodes []Node
	for _, st := range generatedStates {
		newNodes = append(newNodes, *NewNode(node.Depth+1, st, node.costGenerator, node))
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
