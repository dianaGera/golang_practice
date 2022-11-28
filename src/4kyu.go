package challange

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Letters struct {
	count int
	set   string
	char  string
}

// NOT Solved
//
//	It("should handle basic cases 1", func() {
//		dotest("Are they here", "yes, they are here", "2:eeeee/2:yy/=:hh/=:rr")
//		dotest("uuuuuu", "uuuuuu", "=:uuuuuu")
//		dotest("looping is fun but dangerous", "less dangerous than coding",
//				"1:ooo/1:uuu/2:sss/=:nnn/1:ii/2:aa/2:dd/2:ee/=:gg")
func Mix() {
	// create dict key -> letter {s1: count, s2: count}
	// create list[]Letters with [count, s num, char]
	// do sort
	s1, s2 := "Are they here", "yes, they are here"

	charset := make(map[string]map[int]int)
	for _, v := range s1 {

		if string(v) == strings.ToLower(string(v)) && unicode.IsLetter(v) {
			if charset[string(v)][1] == 0 {
				charset[string(v)] = map[int]int{1: 1}
			} else {
				charset[string(v)][1]++
			}
		}
	}
	for _, v := range s2 {

		if string(v) == strings.ToLower(string(v)) && unicode.IsLetter(v) {
			if charset[string(v)][2] == 0 {
				charset[string(v)] = map[int]int{1: charset[string(v)][1], 2: 1}
			} else {
				charset[string(v)][2]++
			}
		}
	}
	res_list := []Letters{}
	for letter, val := range charset {
		temp := "="
		if val[1] > val[2] && val[1] > 1 {
			temp = "1"
		} else if val[2] > val[1] && val[2] > 1 {
			temp = "2"
		}
		set, _ := strconv.Atoi(temp)
		if set == 0 {
			set = val[1]
		}

		if set > 1 {
			res_list = append(
				res_list,
				Letters{
					count: val[set],
					set:   temp,
					char:  letter,
				},
			)
		}
	}

	// buble sort
	for i, _ := range res_list {
		for j, _ := range res_list[0 : len(res_list)-i-1] {

			if res_list[j].count < res_list[j+1].count {
				temp := res_list[j].count
				res_list[j].count = res_list[j-1].count
				res_list[j-1].count = temp
			}

		}
	}
	fmt.Println(sort.StringsAreSorted(res_list))

}
