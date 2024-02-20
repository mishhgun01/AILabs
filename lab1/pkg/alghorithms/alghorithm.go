package alghorithms

import (
	"AILabs/lab1/pkg/consts"
	"AILabs/lab1/pkg/state"
	"fmt"
)

// DFS - Алгоритм поиска в глубину с генерацией состояний.
func DFS(startState *state.State) []state.State {
	chain, _ := search(consts.INFINITY, startState)
	return chain
}

// IDDFS - алгоритм поиска в глубину с итеративным углублением.
func IDDFS(startState *state.State) []state.State {

	var queue []state.State
	depth := 1

	queue = append(queue, *startState)
	for depth < consts.INFINITY {
		state := queue[0]
		fmt.Println(queue)
		for len(queue) > 0 {
			chain, ok := search(depth, startState)
			if ok {
				return chain
			}
			queue = queue[1:]
		}

		substates := state.GenerateSubstates()
		queue = append(queue, substates...)
		depth++
	}
	return nil
}

// search - функция поиска в глубину.
func search(maxDepth int, startState *state.State) ([]state.State, bool) {
	var stack []state.State
	var chain []state.State
	stack = append(stack, *startState)

	depth := 0
	for depth < maxDepth && len(stack) > 0 {
		// Получаем текущее состояние из стека.
		state := stack[len(stack)-1]
		// Убираем текущее состояние из стека.
		stack = stack[:len(stack)-1]

		// Если текущее состояние является конечным, возвращаем всю цепочку.
		if state.IsResult() {
			return chain, true
		}

		// Генерация последовательностей.
		substates := state.GenerateSubstates()
		if len(substates) == 0 {
			// Если состояний больше нет - значит мы дошли до листа и не нашли ответ. Убираем лист из цепи.
			chain = nil
			continue
		}

		stack = append(stack, substates...)
		depth++
	}

	return nil, false
}
