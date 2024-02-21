package node

// необходимо больше для второй лабы
type PathCostFunction func(node *Node) int

func PathCostDFS(node *Node) int {
	if node == nil {
		return 1
	}
	return node.Depth
}
