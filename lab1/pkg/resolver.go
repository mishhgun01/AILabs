package pkg

// resolver - шаблон функции решения задачи
type resolver func(position) []state

// Resolve - функция решающая задачу.
func Resolve(state *state, generate resolver) []state {
	return nil
}
