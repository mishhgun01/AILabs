package node

// необходимо больше для второй лабы
type pathCostGenerator func(node *Node) int

func DefaultDepthGenerator(node *Node) int {
	if node == nil {
		return 1
	}
	return node.Depth
}
