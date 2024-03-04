package solver

import (
	"AILabs/lab1/pkg/node"
)

// alghorithm - шаблонная функция алгоритма, решающего задачу.
type alghorithm func(*node.Node) ([]node.Node, int, int)

// Resolve - функция решающая задачу.
func Resolve(st *node.Node, f alghorithm) ([]node.Node, int, int) {
	return f(st)
}
