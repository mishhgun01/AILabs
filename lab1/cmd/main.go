package main

import (
	"AILabs/lab1/pkg/alghorithms"
	"AILabs/lab1/pkg/consts"
	"AILabs/lab1/pkg/node"
	"AILabs/lab1/pkg/solver"
	"AILabs/lab1/pkg/state"
	"fmt"
	"os"
	"time"
)

func main() {

	var algoChosen, stepsChosen int
	var withSteps bool
	fmt.Println("Выберите алгоритм:\n1 - В глубину\n2 - В глубину с итеративным углублением")
	fmt.Print("->")
	fmt.Scan(&algoChosen)
	fmt.Println("Пошагово или автоматически:\n1 - Пошагово\n2 - Автоматически")
	fmt.Print("->")
	fmt.Scan(&stepsChosen)
	if stepsChosen == 1 {
		withSteps = true
	}
	file, err := os.OpenFile(consts.FILENAME, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var algo func(startNode *node.Node, option bool) ([]node.Node, int, int)

	switch algoChosen {
	case 1:
		algo = alghorithms.DFS
	case 2:
		algo = alghorithms.IDDFS
	default:
		fmt.Println("Неизвестный алгоритм")
		return
	}

	startNode := node.NewNode(1, state.StartState(), node.DefaultDepthGenerator, nil)
	start := time.Now()
	output, steps, nodeCount := solver.Resolve(startNode, withSteps, algo)
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
