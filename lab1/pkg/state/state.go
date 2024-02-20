package state

// State - модель состояния.
type State struct {
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

	CheckedStates []State
)

// StartState - инициализация начального состояния.
func StartState() *State {
	return &State{
		matrix: startStateMatrix,
	}
}
