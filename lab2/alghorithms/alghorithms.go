package alghorithms

import (
	"AILabs/lab1/pkg/node"
	"fmt"
	"os"
)

// search - функция поиска.
func Search(startNode *node.Node, option bool) ([]node.Node, int, int) {
	var queue []node.Node
	nodeCount := 1
	steps := 0

	//var chain []node.Node
	queue = append(queue, *startNode)
	for len(queue) > 0 {
		// Получаем текущее состояние из стека.
		if nodeCount < len(queue) {
			nodeCount = len(queue)
		}
		currentNode := queue[0]
		// Убираем текущее состояние из стека.
		queue = queue[1:]

		// Если текущее состояние является конечным, возвращаем всю цепочку.
		if currentNode.State.IsResult() {
			return currentNode.ChainSlice(), steps, nodeCount
		}
		// Генерация последовательностей.
		subnodes := currentNode.GenerateSubnodes()
		steps++
		queue = append(queue, subnodes...)
		queue = node.SortByPathCost(queue)
		// Если состояний больше нет - значит мы дошли до листа и не нашли ответ. Убираем лист из цепи.
		if option {
			fmt.Println("Шаг:", steps)
			fmt.Println("Текущее состояние:")
			currentNode.Print(os.Stdout)
			fmt.Fprintf(os.Stdout, "Является ли решением:%t \n", currentNode.State.IsResult())
			fmt.Println("Сгенерированные состояния:")
			for _, item := range subnodes {
				item.Print(os.Stdout)
			}
			fmt.Println("Состояния ожидающие раскрытия:")
			for _, item := range queue {
				item.Print(os.Stdout)
			}
			fmt.Scanln()
		}
	}

	return nil, steps, nodeCount
}
