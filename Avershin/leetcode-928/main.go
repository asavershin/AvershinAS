package main

type Set map[int]struct{}

func NewSet(len int) Set {
	return make(Set, len)
}

func (s Set) Add(element int) {
	s[element] = struct{}{}
}

func (s Set) Remove(element int) {
	delete(s, element)
}

func (s Set) Contains(element int) bool {
	_, exists := s[element]
	return exists
}

func (s Set) Size() int {
	return len(s)
}

func minMalwareSpread(graph [][]int, initial []int) (res int) {
	res = 301
	if len(initial) == 1 {
		return initial[0]
	}

	adjGr := make([][]int, len(graph))
	for i := 0; i < len(graph); i++ {
		for j := i + 1; j < len(graph); j++ {
			if graph[i][j] == 1 {
				adjGr[i] = append(adjGr[i], j)
				adjGr[j] = append(adjGr[j], i)
			}
		}
	}
	var initMap = make(map[int]int)
	for _, i := range initial {
		initMap[i] = 0
	}
	var infFrom = make(map[int][]int, len(graph)-len(initial))
	for i := range len(graph) {
		if _, ok := initMap[i]; !ok {
			infFrom[i] = make([]int, 0)
		}
	}
	for key := range infFrom {
		dfs(adjGr, key, key, initMap, NewSet(len(graph)), infFrom)
	}
	for _, v := range infFrom {
		if len(v) == 1 {
			initMap[v[0]]++
		}
	}
	var tmp int
	var v int
	for _, i := range initial {
		v = initMap[i]
		if v > tmp {
			res = i
			tmp = v
		} else if tmp == v && i < res {
			res = i
		}
	}
	return res
}

func dfs(adjGr [][]int, root int, from int, initial map[int]int, visited Set, infFrom map[int][]int) {
	visited.Add(from)
	for _, to := range adjGr[from] {
		if len(infFrom[root]) > 1 {
			return
		}
		if visited.Contains(to) {
			continue
		}
		if _, ok := initial[to]; ok {
			visited.Add(to)
			infFrom[root] = append(infFrom[root], to)
			continue
		}
		dfs(adjGr, root, to, initial, visited, infFrom)
	}
}

func main() {

}
