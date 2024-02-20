package main

import (
	"AILabs/lab1/pkg/alghorithms"
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

	file, err := os.OpenFile(fname, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	startState := state.StartState()

	t1 := time.Now().UnixNano()
	output := solver.Resolve(startState, alghorithms.IDDFS)

	t2 := time.Now().UnixNano()
	for _, state := range output {
		state.Print(file)
	}

	fmt.Fprintf(file, "Прошло %d секунд\n", t2-t1/1000000000)
}
