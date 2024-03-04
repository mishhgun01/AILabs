package node

import "sort"

// необходимо больше для второй лабы
type pathCostGenerator func(node *Node)

func DefaultDepthGenerator(node *Node) {}

func MinimumStepsDepthGenerator(node *Node) {
	node.pathCost = node.State.GetStepsValue()
}

func ManhattanDistanceGenerator(node *Node) {
	node.pathCost = node.State.GetManhattanDistance()
}

func SortByPathCost(nodes []Node) []Node {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].pathCost < nodes[j].pathCost
	})

	return nodes
}

func AStarMinimumStepsDepthGenerator(node *Node) {
	node.pathCost = node.Depth + node.State.GetStepsValue()
}

func AStarManhattanDistanceGenerator(node *Node) {
	node.pathCost = node.Depth + node.State.GetManhattanDistance()
}
