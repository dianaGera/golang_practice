package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Letters struct {
	count int
	set   byte
	char  byte
}

func main() {
	s1, s2 := "looping is fun but dangerous", "less dangerous than coding"

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
	for i, v := range charset {
		fmt.Println(i, v)
	}
	res_list := []Letters{}
	for letter, val := range charset {
		// fmt.Println(letter)
		temp := "3"
		set := 0
		if val[1] > val[2] {
			temp = "1"
			set = 1
		} else if val[2] > val[1] {
			temp = "2"
			set = 2
		} else {
			temp = "3"
			set = 1
		}

		if val[set] > 1 {
			res_list = append(
				res_list,
				Letters{
					count: val[set],
					set:   temp[0],
					char:  letter[0],
				},
			)
		}
	}

	for i := 0; i < len(res_list)-1; i++ {
		for j := 0; j < len(res_list)-i-1; j++ {

			if int(res_list[j].char) > int(res_list[j+1].char) {
				res_list[j], res_list[j+1] = res_list[j+1], res_list[j]
			}
			if byte(res_list[j].set) > byte(res_list[j+1].set) {
				res_list[j], res_list[j+1] = res_list[j+1], res_list[j]
			}
			if res_list[j].count < res_list[j+1].count {
				res_list[j], res_list[j+1] = res_list[j+1], res_list[j]
			}
		}
	}
	res_str := ""
	for _, v := range res_list {
		if string(v.set) != "3" {
			res_str += string(v.set) + ":"
		} else {
			res_str += "=:"
		}
		res_str += strings.Repeat(string(v.char), v.count) + "/"

	}
	fmt.Println(res_str[:len(res_str)-1])
}
