package main

import (
	"AILabs/lab1/pkg"
)

func main() {
	startState := pkg.StartState()
	startState.Print()
	output := pkg.Resolve(startState, pkg.DFS)
	for _, state := range output {
		state.Print()
	}
}
