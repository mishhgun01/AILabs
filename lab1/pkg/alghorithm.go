package pkg

type position string

const (
	LEFT   = "Left"
	RIGHT  = "Right"
	TOP    = "Top"
	BOTTOM = "Bottom"
)

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
// TODO
func (state *state) generateSubstates(swapPosition position) []state {
	switch swapPosition {
	case LEFT:
	case RIGHT:
	case TOP:
	case BOTTOM:
	}
	return nil
}
