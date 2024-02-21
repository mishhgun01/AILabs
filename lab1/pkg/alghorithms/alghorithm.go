package alghorithms

import (
	"AILabs/lab1/pkg/consts"
	"AILabs/lab1/pkg/node"
	"AILabs/lab1/pkg/state"
	"fmt"
)

// DFS - Алгоритм поиска в глубину с генерацией состояний.
func DFS(startNode *node.Node) []node.Node {
	chain, _ := search(consts.INFINITY, startNode)
	return chain
}

// IDDFS - алгоритм поиска в глубину с итеративным углублением
func IDDFS(startNode *node.Node) []node.Node {
	depth := 1
	for depth < consts.INFINITY {

		chain, ok := search(depth, startNode)
		if ok {
			return chain
		}
		state.CheckedStates = nil

		depth++
		//time.Sleep(time.Nanosecond)

	}

	return nil
}

// IDDFS - алгоритм поиска в глубину с итеративным углублением

// search - функция поиска в глубину.
func search(maxDepth int, startNode *node.Node) ([]node.Node, bool) {
	var stack []node.Node
	//var chain []node.Node
	stack = append(stack, *startNode)
	for len(stack) > 0 {
		// Получаем текущее состояние из стека.
		vertex := stack[len(stack)-1]
		// Убираем текущее состояние из стека.
		stack = stack[:len(stack)-1]

		// Если текущее состояние является конечным, возвращаем всю цепочку.
		if vertex.State.IsResult() {
			fmt.Println(vertex.Depth)
			return vertex.GetChain(), true
		}

		// Генерация последовательностей.

		if vertex.Depth < maxDepth {
			subnodes := vertex.GenerateSubNodes(node.PathCostDFS)

			stack = append(stack, subnodes...)
			// Если состояний больше нет - значит мы дошли до листа и не нашли ответ. Убираем лист из цепи.
			continue
		}

	}

	return nil, false
}
