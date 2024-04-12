package main

import "fmt"

func findChampion(n int, edges [][]int) (champion int) {
	var flagIn = make([]int, n)
	for _, edge := range edges {
		flagIn[edge[1]]++
	}
	champion = -1
	for cmd, i := range flagIn {
		if i == 0 {
			if champion != -1 {
				return -1
			}
			champion = cmd
		}
	}
	return champion
}

func main() {
	fmt.Println(findChampion(3, [][]int{{0, 1}, {1, 2}}))
}
