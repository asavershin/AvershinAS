package main

import (
	"container/list"
)

func createGraph(dislikes [][]int, graph map[int]*list.List) {
	for _, dislike := range dislikes {
		vertex1 := dislike[0]
		vertex2 := dislike[1]

		if _, ok := graph[vertex1]; !ok {
			graph[vertex1] = list.New()
		}

		graph[vertex1].PushBack(vertex2)
	}
}

func possibleBipartition(n int, dislikes [][]int) bool {
	if len(dislikes) == 0 {
		return true
	}
	graph := make(map[int]*list.List, n+1)
	createGraph(dislikes, graph)
	color := make(map[int]bool)

	stack := list.New()
	stack.PushFront(1)
	color[1] = true // Начинаем с красного цвета

	for stack.Len() > 0 {
		node := stack.Remove(stack.Front()).(int)

		for adj := graph[node].Front(); adj != nil; adj = adj.Next() {
			neighbor := adj.Value.(int)

			// Проверяем, если у соседа уже установлен цвет
			if c, exists := color[neighbor]; exists {
				// Если цвет соседа совпадает с цветом текущего узла, то это не двудольный граф
				if c == color[node] {
					return false
				}
			} else {
				// Устанавливаем цвет соседа
				color[neighbor] = !color[node]
				stack.PushFront(neighbor)
			}
		}
	}

	return true
}

func main() {

}
