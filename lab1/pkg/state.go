package pkg

import "reflect"

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

func CompareState(st1, st2 state) bool {
	if !reflect.DeepEqual(st1.matrix, st2.matrix) {
		return false
	}
	return true
}

// получение всех возможных состояний из текущего
func (st *state) GetNewState() []state {
	return nil
}
