package main

import "fmt"

func findSmallestSetOfVertices(n int, edges [][]int) []int {
    origins := make([]int, n)
    for _, edge := range edges {
        origins[edge[1]]++
    }
    
    result := []int{}
    for vertex, frequency := range origins {
        if frequency == 0 {
            result = append(result, vertex)
        }
    }
    
    return result
}

func main() {
    edges1 := [][]int{{0,1},{0,2},{2,5},{3,4},{4,2}}
    fmt.Println(findSmallestSetOfVertices(6, edges1))
    
    edges2 := [][]int{{0,1},{2,1},{3,1},{1,4},{2,4}}
    fmt.Println(findSmallestSetOfVertices(5, edges2)) 
}