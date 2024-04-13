package main

func intToOneOrZero(to int) int {
	if to > 0 {
		return 1
	}
	return 0
}

func abs(to int) int {
	if to < 0 {
		return -to
	}
	return to
}
func bfs(graph *[100000][]int, from int, visited *[100000]bool) (res int) {
	var st = make([]int, 0)
	st = append(st, from)
	for len(st) > 0 {
		from, st = st[len(st)-1], st[:len(st)-1]
		visited[from] = true
		for _, to := range graph[from] {
			if !visited[abs(to)] {
				st = append(st, abs(to))
				res += intToOneOrZero(to)
			}
		}
	}
	return res
}
func minReorder(n int, connections [][]int) int {
	var graph = [100000][]int{}
	for _, c := range connections {
		graph[c[0]] = append(graph[c[0]], c[1])
		graph[c[1]] = append(graph[c[1]], -c[0])
	}
	return bfs(&graph, 0, &[100000]bool{false})
}

func main() {

}
