package main

import (
	"fmt"
)

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func main() {
	res := PosPeaks{}
	// {Pos: [1, 5], Peaks: [12, 14]}
	// array := []int{-3, 12, 12, 9, -3, 14, 5, 1, -4}
	array := []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1}
	heap := [4]int{0, 0, 0, 0}
	for i, v := range array[1:len(array)-1] {
		if heap[1] != array[i] {
			heap[0] = array[i]
		} 
		if array[i+2] == v {
			heap[2] += 1
		} else {
			heap[3] = heap[2]
			heap[2] = 0
		}
		
		heap[1] = array[i+2]
		
		if heap[0] < v && v > heap[1] {
			res.Pos = append(res.Pos, i-heap[3]+1)
			res.Peaks = append(res.Peaks, v) 
		}
		
	}
	fmt.Println(res)
	
	
}
