package pkg

type position string

// DFS - Алгоритм поиска в глубину с генерацией состояний.
func DFS(startState *state) []state {
	var stack []state
	var chain []state
	stack = append(stack, *startState)

	for len(stack) > 0 {
		// Получаем текущее состояние из стека.
		state := stack[len(stack)-1]
		// Убираем текущее состояние из стека.
		stack = stack[:len(stack)-1]
		chain = append(chain, state)

		// Если текущее состояние является конечным, возвращаем всю цепочку.
		if state.isResult() {
			return chain
		}

		// Генерация последовательностей.
		substates := state.generateSubstates()
		if len(substates) == 0 {
			// Если состояний больше нет - значит мы дошли до листа и не нашли ответ. Чистим цепочку.
			chain = chain[:]
			continue
		}

		stack = append(stack, substates...)

	}
	return nil
}
