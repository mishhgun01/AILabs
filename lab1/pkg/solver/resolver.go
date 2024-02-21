package solver

import (
	"AILabs/lab1/pkg/node"
)

// alghorithm - шаблонная функция алгоритма, решающего задачу.
type alghorithm func(*node.Node) []node.Node

// Resolve - функция решающая задачу.
func Resolve(st *node.Node, f alghorithm) []node.Node {
	return f(st)
}
