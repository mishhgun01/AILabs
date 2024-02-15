package main

import (
	"AILabs/lab1/pkg"
	"fmt"
)

func main() {
	startState := pkg.StartState()
	chain := pkg.Resolve(startState, pkg.Alghorithm)
	fmt.Println(chain)
}
