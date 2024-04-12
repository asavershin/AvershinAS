package main

import (
	"container/list"
	"fmt"
)

type VertEdge struct {
	vert string
	edge float64
}

func bfs(g map[string][]VertEdge, from, to string) float64 {

	visited := make(map[string]bool)
	dist := make(map[string]float64)
	queue := list.New()
	queue.PushBack(from)
	visited[from] = true
	dist[from] = 1

	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(string)

		for _, next := range g[curr] {
			if !visited[next.vert] {
				dist[next.vert] = dist[curr] * next.edge
				visited[next.vert] = true
				queue.PushBack(next.vert)
			}
		}
	}
	res := dist[to]
	if res == 0 {
		return -1
	}
	return res
}

func check(equations [][]string, querie []string, res *[]float64) bool {
	flag1 := true
	flag2 := true
	for _, equation := range equations {
		if equation[0] == querie[0] || equation[1] == querie[0] {
			flag1 = false
		}
		if equation[0] == querie[1] || equation[1] == querie[1] {
			flag2 = false
		}
		if flag1 == false && flag2 == false {
			return false
		}
	}
	*res = append(*res, -1)
	return true
}

func calcEquation(equations [][]string, values []float64, queries [][]string) (res []float64) {
	var edjes = make(map[string][]VertEdge)

	for i, equation := range equations {
		edjes[equation[0]] = append(edjes[equation[0]], VertEdge{vert: equation[1], edge: values[i]})
		edjes[equation[1]] = append(edjes[equation[1]], VertEdge{vert: equation[0], edge: 1. / values[i]})
	}

	for _, querie := range queries {

		if check(equations, querie, &res) {
			continue
		}

		if querie[0] == querie[1] {
			res = append(res, 1)
			continue
		}
		res = append(res, bfs(edjes, querie[0], querie[1]))
	}
	return
}

func main() {

	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println(calcEquation(equations, values, queries))
}
