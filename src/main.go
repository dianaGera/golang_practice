package main

import (
	"fmt"
	"strings"
)

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func main() {
	str := "stress"
	char_dict := map[string][]int{}
	res := -1

	for i, v := range str {
		ch := char_dict[strings.ToUpper(string(v))]
		if len(ch) == 0 {
			char_dict[strings.ToUpper(string(v))] = []int{i, 1}
		} else {
			char_dict[strings.ToUpper(string(v))][1]++
		}
	}
	for _, val := range char_dict {
		if val[1] == 1 {
			if res == -1 || res > val[0] {
				res = val[0]
			}
		}
	}
	fmt.Println(string(str[res]))

}
