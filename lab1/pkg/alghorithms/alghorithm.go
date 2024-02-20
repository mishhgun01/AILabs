package alghorithms

import (
	"AILabs/lab1/pkg/consts"
	"AILabs/lab1/pkg/state"
)

// DFS - Алгоритм поиска в глубину с генерацией состояний.
func DFS(startState *state.State) []state.State {
	chain, _ := search(consts.INFINITY, startState)
	return chain
}

// IDDFS - алгоритм поиска в глубину с итеративным углублением
func IDDFS(startState *state.State) []state.State {
	depth := 1
	for depth < consts.INFINITY {

		chain, ok := search(depth, startState)
		if ok {
			return chain
		}
		depth++
		//time.Sleep(time.Nanosecond)
		state.CheckedStates = state.CheckedStates[:0]
	}

	return nil
}

// IDDFS - алгоритм поиска в глубину с итеративным углублением

// search - функция поиска в глубину.
func search(maxDepth int, startState *state.State) ([]state.State, bool) {
	var stack []state.State
	var chain []state.State
	stack = append(stack, *startState)
	depth := 0
	for depth < maxDepth && len(stack) > 0 {
		// Получаем текущее состояние из стека.
		state := stack[len(stack)-1]
		chain = append(chain, state)
		// Убираем текущее состояние из стека.
		stack = stack[:len(stack)-1]

		// Если текущее состояние является конечным, возвращаем всю цепочку.
		if state.IsResult() {
			return chain, true
		}

		// Генерация последовательностей.
		substates := state.GenerateSubstates()
		if len(substates) == 0 {
			depth--
			// Если состояний больше нет - значит мы дошли до листа и не нашли ответ. Убираем лист из цепи.
			chain = chain[:len(chain)-1]
			continue
		}

		stack = append(stack, substates...)
		depth++
	}

	return nil, false
}
