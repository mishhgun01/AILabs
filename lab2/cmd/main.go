package main

import (
	"AILabs/lab1/pkg/consts"
	"AILabs/lab1/pkg/node"
	"AILabs/lab1/pkg/solver"
	"AILabs/lab1/pkg/state"
	"AILabs/lab2/alghorithms"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile(consts.FILENAME, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	startNode := node.NewNode(1, state.StartState(), node.MinimumStepsDepthGenerator, nil)
	start := time.Now()
	output, steps, nodeCount := solver.Resolve(startNode, alghorithms.AStar)

	fmt.Println("STeps:", steps)
	fmt.Println("Max nodes:", nodeCount)

	duration := time.Since(start)
	for _, vertex := range output {
		vertex.State.Print(file)
	}

	fmt.Fprintf(file, "Прошло %s\n", duration)
}
