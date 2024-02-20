package alghorithms

import (
	"AILabs/lab1/pkg/state"
)

// DFS - Алгоритм поиска в глубину с генерацией состояний.
func DFS(startState *state.State) []state.State {
	var stack []state.State
	var chain []state.State
	stack = append(stack, *startState)

	for len(stack) > 0 {
		// Получаем текущее состояние из стека.
		state := stack[len(stack)-1]
		// Убираем текущее состояние из стека.
		stack = stack[:len(stack)-1]
		chain = append(chain, state)

		// Если текущее состояние является конечным, возвращаем всю цепочку.
		if state.IsResult() {
			return chain
		}

		// Генерация последовательностей.
		substates := state.GenerateSubstates()
		if len(substates) == 0 {
			// Если состояний больше нет - значит мы дошли до листа и не нашли ответ. Убираем лист из цепи.
			chain = chain[:len(chain)-1]
			continue
		}

		stack = append(stack, substates...)

	}
	return nil
}
