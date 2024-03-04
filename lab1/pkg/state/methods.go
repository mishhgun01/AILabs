package state

import (
	"AILabs/lab1/pkg/consts"
	"fmt"
	"math"
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
	CheckedStates = append(CheckedStates, *st)
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
				if !state.checked() {
					newStates = append(newStates, state)
				}
			}

			return newStates
		}
	}
	return nil
}

// checked - Проверка состояние на то, что оно было пройдено.
func (st *State) checked() bool {
	for _, item := range CheckedStates {
		if item.matrix == st.matrix {
			return true
		}
	}

	return false
}

func (st *State) GetStepsValue() int {
	var steps int
	for i := range st.matrix {
		for j := range st.matrix[i] {
			if st.matrix[i][j] != endStateMatrix[i][j] {
				steps++
			}
		}
	}
	return steps
}

func (st *State) GetManhattanDistance() int {
	var distance int
	endStateIndexes := make(map[int][2]int)
	for i := range endStateMatrix {
		for j := range endStateMatrix[i] {
			endStateIndexes[endStateMatrix[i][j]] = [2]int{i, j}
		}
	}

	for i := range st.matrix {
		for j := range st.matrix[i] {
			xEnd, yEnd := endStateIndexes[st.matrix[i][j]][0], endStateIndexes[st.matrix[i][j]][1]
			xStart, yStart := i, j
			distance += int(math.Abs(float64(xStart-xEnd)) + math.Abs(float64(yStart-yEnd)))
		}
	}

	return distance

}
