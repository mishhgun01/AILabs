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
	var algoChosen, stepsChosen int
	var withSteps bool
	fmt.Println("Выберите алгоритм:\n1 - Жадный с h1\n2 - Жадный с h2\n3 - А* с h1\n4 - А* с h2")
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

	var algo func(*node.Node)

	defer file.Close()
	switch algoChosen {
	case 1:
		algo = node.MinimumStepsDepthGenerator
	case 2:
		algo = node.ManhattanDistanceGenerator
	case 3:
		algo = node.AStarMinimumStepsDepthGenerator
	case 4:
		algo = node.AStarManhattanDistanceGenerator
	default:
		fmt.Println("Неизвестный алгоритм")
		return
	}

	startNode := node.NewNode(1, state.StartState(), algo, nil)
	start := time.Now()
	output, steps, nodeCount := solver.Resolve(startNode, withSteps, alghorithms.Search)

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
