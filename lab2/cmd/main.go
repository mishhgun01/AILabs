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

	startNode := node.NewNode(1, state.StartState(), node.AStarMinimumStepsDepthGenerator, nil)
	start := time.Now()
	output, steps, nodeCount := solver.Resolve(startNode, true, alghorithms.AStar)

	duration := time.Since(start)
	if output != nil {
		for _, vertex := range output {
			vertex.State.Print(file)
		}
	} else {
		fmt.Fprintf(file, "Решений нет\n")
	}

	fmt.Fprintf(file, "Прошло %s\n", duration)
	fmt.Fprintf(file, "Шаги %d\n", steps)
	fmt.Fprintf(file, "Макс количество вершин в стеке %d\n", nodeCount)
	fmt.Fprintf(file, "Глубина: %d\n", len(output))
}
