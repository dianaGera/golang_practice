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
	walk := []rune{'e', 'w', 'w', 'w', 'e', 'e', 'w', 'e'}
	res := [2]int{0, 0}

	for _, v := range walk {
		println(v)
		switch v {
			case 110: res[0]++;
			case 115: res[0]--;
			case 119: res[1]++;
			case 101: res[1]--;
		}
	}
	fmt.Println(res)
	// if res[0] + res[1] == 0 {
	// 	return true
	// }
}
