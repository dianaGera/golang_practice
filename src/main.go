package main

import (
	"fmt"
	// 'strings'
)

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func main() {
	arr := []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}
	dir := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"WEST": "EAST",
		"EAST": "WEST",
	}
	res := []string{"START"}
	for _, v := range arr {
		if res[len(res)-1] == dir[v] {
			res = res[0:len(res)-1]
		} else {
			res = append(res, v)
		}

		fmt.Println(res[1:])
	}
	
}
