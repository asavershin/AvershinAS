package main

import (
	"container/list"
	"fmt"
)

func square(x int) int {
	return x * x
}

type graph map[int][]int

func makeGraph(bombs *[][]int) *graph {
	g := make(graph)
	for i, bomb1 := range *bombs {
		for k, bomb2 := range *bombs {
			if i == k {
				continue
			}
			if square(bomb1[0]-bomb2[0])+square(bomb1[1]-bomb2[1]) <= square(bomb1[2]) {
				g[i] = append(g[i], k)
			}
		}
	}
	return &g
}

func dfs(g *graph, start int) (res int) {

	st := list.New()

	visited := make(map[int]bool, len(*g))
	st.PushFront(start)
	visited[start] = true

	for st.Len() > 0 {
		curr := st.Remove(st.Front()).(int)

		res++

		for _, next := range (*g)[curr] {
			if !visited[next] {
				st.PushFront(next)
				visited[next] = true
			}
		}
	}

	return
}

func maximumDetonation(bombs [][]int) (res int) {
	g := makeGraph(&bombs)
	if len(*g) == 0 {
		return 1
	}
	for curr := range *g {
		tmp := dfs(g, curr)
		if tmp > res {
			res = tmp
		}
	}
	return
}

func main() {
	bombs := [][]int{{2, 1, 3}, {6, 1, 4}}
	fmt.Println(maximumDetonation(bombs))
}
