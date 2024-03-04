package alghorithms

import (
	"AILabs/lab1/pkg/consts"
	"AILabs/lab1/pkg/node"
	"os"
)

func AStar(startNode *node.Node) ([]node.Node, int, int) {
	chain, steps, nodeCount, _ := search(consts.INFINITY, startNode)
	return chain, steps, nodeCount
}

// search - функция поиска.
func search(maxDepth int, startNode *node.Node) ([]node.Node, int, int, bool) {
	var queue []node.Node
	nodeCount := 1
	steps := 0

	//var chain []node.Node
	queue = append(queue, *startNode)
	for len(queue) > 0 {
		// Получаем текущее состояние из стека.
		if nodeCount < len(queue) {
			nodeCount = len(queue)
		}
		currentNode := queue[0]
		// Убираем текущее состояние из стека.
		queue = queue[1:]

		// Если текущее состояние является конечным, возвращаем всю цепочку.
		if currentNode.State.IsResult() {
			return currentNode.ChainSlice(), steps, nodeCount, true
		}

		currentNode.State.Print(os.Stdout)

		// Генерация последовательностей.
		subnodes := currentNode.GenerateSubnodes()
		steps++
		queue = append(queue, subnodes...)
		queue = node.SortByPathCost(queue)
		// Если состояний больше нет - значит мы дошли до листа и не нашли ответ. Убираем лист из цепи.

	}

	return nil, steps, nodeCount, false
}
