package pkg

const (
	LEFT  = "Left"
	RIGHT = "Right"
	UP    = "Up"
	DOWN  = "Down"
)

// generateSubstates - генерация последовательностей.
func (st *state) generateSubstates() []state {
	var generatedStates []state
	// Сразу добавляем текущее состояние в провереннные.
	checkedStates = append(checkedStates, *st)
	for i := range st.matrix {
		for j := range st.matrix[i] {
			if st.matrix[i][j] != EMPTY_ELEMENT {
				continue
			}
			// Когда находим пустой элемент, двигаем его в разные стороны.
			matrixLeft := swap(i, j, LEFT, st.matrix)
			matrixRight := swap(i, j, RIGHT, st.matrix)
			matrixTop := swap(i, j, UP, st.matrix)
			matrixBottom := swap(i, j, DOWN, st.matrix)

			// Добавляем в сгенерированные состояния новые состояния.
			generatedStates = append(generatedStates,
				state{
					matrix: matrixLeft,
				},
				state{
					matrix: matrixRight,
				},
				state{
					matrix: matrixTop,
				},
				state{
					matrix: matrixBottom,
				},
			)

			// Проверяем состояния на то, что они были пройдены ранее.
			var newStates []state
			for _, state := range generatedStates {
				if !state.stateChecked() {
					newStates = append(newStates, state)
				}
			}

			return newStates
		}
	}
	return nil
}

// swap - Перестановка элементов матрицы.
func swap(i, j int, pos position, oldMatrix [3][3]int) [3][3]int {
	matrix := oldMatrix
	switch pos {
	case LEFT:
		if j == 0 {
			return oldMatrix
		}
		matrix[i][j-1], matrix[i][j] = matrix[i][j], matrix[i][j-1]
		return matrix
	case RIGHT:
		if j == 2 {
			return oldMatrix
		}
		matrix[i][j+1], matrix[i][j] = matrix[i][j], matrix[i][j+1]
		return matrix
	case UP:
		if i == 0 {
			return oldMatrix
		}
		matrix[i][j], matrix[i-1][j] = matrix[i-1][j], matrix[i][j]
		return matrix
	case DOWN:
		if i == 2 {
			return oldMatrix
		}
		matrix[i][j], matrix[i+1][j] = matrix[i+1][j], matrix[i][j]
		return matrix
	}

	return matrix
}
