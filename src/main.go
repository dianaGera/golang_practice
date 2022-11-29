package main

import (
	"fmt"
	"strings"
)

type Letters struct {
	count int
	idx   int
}

func main() {
	str := "abba"

	if len(str) > 0 {
		str_set := make(map[string]Letters)

		for i, v := range str {
			l := strings.ToLower(string(v))
			if str_set[l].count == 0 {
				str_set[l] = Letters{
					count: 1,
					idx:   i,
				}
			} else {
				entry := str_set[l]
				entry.count++
				str_set[l] = entry
			}
		}
		fmt.Println(str_set)
		idx := len(str)+1
		for key := range str_set {
			if str_set[key].count == 1 {
				if idx > str_set[key].idx {
					idx = str_set[key].idx
				}
			}
		}
		if idx != len(str)+1 {

		}
		fmt.Println(string(str[idx]))
	} else {
		fmt.Println(str)
	}

}
