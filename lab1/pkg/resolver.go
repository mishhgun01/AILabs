package pkg

type alghorithm func(*state) []state

// Resolve - функция решающая задачу.
func Resolve(st *state, f alghorithm) []state {
	return f(st)
}
