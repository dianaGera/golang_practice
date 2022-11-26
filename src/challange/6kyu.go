package challange

import (
	"fmt"
	"strings"
)

func DecodeRomeInt(roman string) int {
	rome_dict := map[string]int{
		  "I": 1,
		  "V": 5, 
		  "X": 10, 
		  "L": 50, 
		  "C": 100, 
		  "D": 500, 
		  "M": 1000,
	  }
  
	  cur_int := string(roman[0])
	  count_sequence := 1
	  res := 0
	  for _, v := range roman[1:] {
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
	return res
  }


func CreatePhoneNumber(numbers [10]uint) string {
	res := []interface{}{}
	for _, val := range numbers {
		res = append(res, val)
	} 
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", res...) 
}

  
func ValidBraces(str1 string) bool {
// "(){}[]"   =>  True
// "([{}])"   =>  True
// "(}"       =>  False

str := []rune{}
str = append(str, rune(str1[0]))
i := 0
	for _, v := range str1[1:] {
		str = append(str, v)
		if len(str) != 1 {
			if int(str[i]/10) == int(v/10) && str[i]< v{
				str = str[:i]
				if i > 0 {
					i -= 1
				}
				
			} else {
				i++
			}
		}
	}
	if len(str) > 0 {
		return false
	}
	return true
}


func evenToUpper(str string) string {
	res := ""
	str = strings.ToUpper(str)
	count := 0
	for _, v := range str {
		if string(v) == " " {count++}
		if count == 0 || count%2 == 0 {
			res += string(v)
		} else {
			res += strings.ToLower(string(v))
		}
		count++
	}

	return res

}


func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
	  return false
	}
	
	x, y := 0, 0
	for _, r := range walk {
	  switch r {
	  case 'n': y++
	  case 's': y--
	  case 'e': x++
	  case 'w': x--
	  }
	}
	
	if x == 0 && y == 0 {
	  return true
	}
	
	return false
  }
  