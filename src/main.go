package main

import (
	"fmt"
)

// I          1
// V          5
// X          10
// L          50
// C          100
// D          500
// M          1,000

func main() {
	rome_dict := map[string]int{
		"I": 1,
		"V": 5, 
		"X": 10, 
		"L": 50, 
		"C": 100, 
		"D": 500, 
		"M": 1000,
	}

	r1 := "MDCLXVI"
	cur_int := string(r1[0])
	count_sequence := 1
	res := 0
	for _, v := range r1[1:] {
		if cur_int == string(v) {
			count_sequence++
		} else {
			if rome_dict[string(cur_int)] > rome_dict[string(v)] {
				res += rome_dict[string(cur_int)] * count_sequence
			} else {
				res -= rome_dict[string(cur_int)] * count_sequence
			}
			cur_int = string(v)
			count_sequence = 1
		}
	}
	res += rome_dict[string(cur_int)] * count_sequence

	fmt.Println(res)
}
