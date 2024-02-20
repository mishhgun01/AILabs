package solver

import (
	"AILabs/lab1/pkg/state"
)

// alghorithm - шаблонная функция алгоритма, решающего задачу.
type alghorithm func(*state.State) []state.State

// Resolve - функция решающая задачу.
func Resolve(st *state.State, f alghorithm) []state.State {
	return f(st)
}
