package main

import (
	"AILabs/lab1/pkg/alghorithms"
	"AILabs/lab1/pkg/solver"
	"AILabs/lab1/pkg/state"
)

func main() {
	startState := state.StartState()
	output := solver.Resolve(startState, alghorithms.DFS)
	for _, state := range output {
		state.Print()
	}
}
