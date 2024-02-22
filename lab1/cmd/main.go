package main

import (
	"AILabs/lab1/pkg/alghorithms"
	"AILabs/lab1/pkg/node"
	"AILabs/lab1/pkg/solver"
	"AILabs/lab1/pkg/state"
	"fmt"
	"os"
	"time"
)

const (
	fname = "output_dfs.txt"
)

func main() {

	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	startNode := node.NewNode(1, state.StartState(), node.DefaultDepthGenerator, nil)
	start := time.Now()
	output := solver.Resolve(startNode, alghorithms.IDDFS)

	duration := time.Since(start)
	for _, vertex := range output {
		vertex.State.Print(file)
	}

	fmt.Fprintf(file, "Прошло %s\n", duration)
}
