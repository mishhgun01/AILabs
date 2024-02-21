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
	fname = "C:\\Users\\mihak\\GolandProjects\\AILabs\\lab1\\cmd\\output_dfs.txt"
)

func main() {

	file, err := os.OpenFile(fname, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	startNode := node.NewNode(1, *state.StartState(), node.PathCostDFS, nil)
	t1 := time.Now().UnixNano()
	output := solver.Resolve(startNode, alghorithms.IDDFS)

	t2 := time.Now().UnixNano()
	for _, vertex := range output {
		vertex.State.Print(file)
	}

	fmt.Fprintf(file, "Прошло %d секунд\n", t2-t1/1000000000)
}
