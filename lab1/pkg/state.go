package pkg

// state - модель состояния.
type state struct {
	matrix [3][3]int
}

// константа для удобства обозначения пустого элемента.
const (
	EMPTY_ELEMENT = 0
)

// Перемннные - начальное состояние, конечное состояние и проверенные состояния.
var (
	startStateMatrix = [3][3]int{
		{EMPTY_ELEMENT, 4, 3},
		{6, 2, 1},
		{7, 5, 8},
	}
	endStateMatrix = [3][3]int{
		{1, 2, 3},
		{4, EMPTY_ELEMENT, 5},
		{6, 7, 8},
	}

	checkedStates []state
)

// StartState - инициализация начального состояния.
func StartState() *state {
	return &state{
		matrix: startStateMatrix,
	}
}

// stateChecked - Проверка состояние на то, что оно было пройдено.
func (state *state) stateChecked() bool {
	for _, item := range checkedStates {
		if item.matrix == state.matrix {
			return true
		}
	}

	return false
}

// isResult - Является ли состояние конечным.
func (state *state) isResult() bool {
	return state.matrix == endStateMatrix
}

// generateSubstates - генерация последовательностей.
func (state *state) generateSubstates() []state {
	return nil
}
