package node

import (
	"AILabs/lab1/pkg/state"
	"fmt"
	"os"
)

type Node struct {
	State state.State
	Depth int

	pathCost      int
	costGenerator func(node *Node)
	parent        *Node
}

func (node *Node) Print(f *os.File) {
	fmt.Fprintf(f, "Глубина:%d, Оценочная стоимость:%d \n", node.Depth, node.pathCost)
	node.State.Print(f)
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
