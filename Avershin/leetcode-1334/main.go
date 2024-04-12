package main

import (
	"fmt"
)

var maxInt = 1000000

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
    array := make([][]int, n)
	for i := range array {
		array[i] = make([]int, n)
	}

	for j := range array{
		for k := range array{
			array[j][k] = maxInt
		}
	}

	for i := range edges {
		array[edges[i][0]][edges[i][1]] = edges[i][2]
		array[edges[i][1]][edges[i][0]] = edges[i][2]
	}

	
	for i := 0; i < n; i++ {
		for j := range array{
			for k := range array{
				if j == k {
					continue
				}

				array[j][k] = min(array[j][k], array[j][i] + array[i][k])

			}
		}
	}
	
	countCities := make([]int, n)
	min := n
	for j := range array{
		for k := range array{
			if j == k {
				continue
			}
			if array[j][k] <= distanceThreshold{
				countCities[j]++
			}
		}
		if countCities[j] <= min{
			min = countCities[j]
		}
	}
	for i := n - 1; i > -1; i--{
		if countCities[i] == min{
			return i
		}
	}
	return 0
}

func main() {
    edges1 := [][]int{{0,1,3977},{2,3,8807},{0,2,2142},{1,3,1201}}
	fmt.Println(findTheCity(4, edges1, 8174))
}