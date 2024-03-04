package alghorithms

import (
	"AILabs/lab1/pkg/consts"
	"AILabs/lab1/pkg/node"
	"AILabs/lab1/pkg/state"
	"fmt"
)

// DFS - Алгоритм поиска в глубину с генерацией состояний.
func DFS(startNode *node.Node) ([]node.Node, int, int) {
	chain, steps, nodeCount, _ := search(consts.INFINITY, startNode)
	return chain, steps, nodeCount
}

// IDDFS - алгоритм поиска в глубину с итеративным углублением
func IDDFS(startNode *node.Node) ([]node.Node, int, int) {
	depth := 1
	steps := 0
	nodeCount := 0
	for depth < consts.INFINITY {

		chain, middleSteps, middleNodeCount, ok := search(depth, startNode)
		steps += middleSteps
		if middleNodeCount > nodeCount {
			nodeCount = middleNodeCount
		}
		if ok {
			return chain, steps, nodeCount
		}
		state.CheckedStates = nil

		depth++
		//time.Sleep(time.Nanosecond)

	}

	return nil, steps, nodeCount
}

// IDDFS - алгоритм поиска в глубину с итеративным углублением

// search - функция поиска в глубину.
func search(maxDepth int, startNode *node.Node) ([]node.Node, int, int, bool) {
	var stack []node.Node
	//var chain []node.Node
	nodeCount := 1
	steps := 0
	stack = append(stack, *startNode)
	for len(stack) > 0 {
		if nodeCount < len(stack) {
			nodeCount = len(stack)
		}
		// Получаем текущее состояние из стека.
		currentNode := stack[len(stack)-1]
		// Убираем текущее состояние из стека.
		stack = stack[:len(stack)-1]

		// Если текущее состояние является конечным, возвращаем всю цепочку.
		if currentNode.State.IsResult() {
			fmt.Println(currentNode.Depth)
			return currentNode.ChainSlice(), steps, nodeCount, true
		}

		// Генерация последовательностей.

		if currentNode.Depth < maxDepth {
			subnodes := currentNode.GenerateSubnodes()
			steps++
			stack = append(stack, subnodes...)
			// Если состояний больше нет - значит мы дошли до листа и не нашли ответ. Убираем лист из цепи.
			continue
		}

	}

	return nil, steps, nodeCount, false
}
