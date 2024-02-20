package state

import (
	"AILabs/lab1/pkg/consts"
	"fmt"
	"os"
)

// Print - красивый вывод состояния.
func (st *State) Print(f *os.File) {
	for i := range st.matrix {
		fmt.Fprintf(f, "| %d %d %d |\n", st.matrix[i][0], st.matrix[i][1], st.matrix[i][2])
	}
	fmt.Fprint(f, "---------------\n")
}

// IsResult - Является ли состояние конечным.
func (st *State) IsResult() bool {
	return st.matrix == endStateMatrix
}

// GenerateSubstates - генерация последовательностей.
func (st *State) GenerateSubstates() []State {
	var generatedStates []State
	// Сразу добавляем текущее состояние в провереннные.
	checkedStates = append(checkedStates, *st)
	for i := range st.matrix {
		for j := range st.matrix[i] {
			if st.matrix[i][j] != EMPTY_ELEMENT {
				continue
			}
			// Когда находим пустой элемент, двигаем его в разные стороны.
			matrixLeft := swap(i, j, consts.LEFT, st.matrix)
			matrixRight := swap(i, j, consts.RIGHT, st.matrix)
			matrixTop := swap(i, j, consts.UP, st.matrix)
			matrixBottom := swap(i, j, consts.DOWN, st.matrix)

			// Добавляем в сгенерированные состояния новые состояния.
			generatedStates = append(generatedStates,
				State{
					matrix: matrixLeft,
				},
				State{
					matrix: matrixRight,
				},
				State{
					matrix: matrixTop,
				},
				State{
					matrix: matrixBottom,
				},
			)

			// Проверяем состояния на то, что они были пройдены ранее.
			var newStates []State
			for _, state := range generatedStates {
				if !state.StateChecked() {
					newStates = append(newStates, state)
				}
			}

			return newStates
		}
	}
	return nil
}

// stateChecked - Проверка состояние на то, что оно было пройдено.
func (st *State) StateChecked() bool {
	for _, item := range checkedStates {
		if item.matrix == st.matrix {
			return true
		}
	}

	return false
}
