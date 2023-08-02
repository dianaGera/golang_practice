// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 
// - - - - - - - - - - Vowel Count - - - - - - - - - - - - - - 
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 

// Return the number (count) of vowels in the given string.
// We will consider a, e, i, o, u as vowels for this Kata (but not y).
// The input string will only consist of lower case letters and/or spaces.


func GetCount(str string) (count int) {
  count = 0
  vowels := "aeiou"
  for _, char := range str {
    for _, vowel := range vowels {
      if char == vowel {
       count++ 
      }
    }
  }
  return count
}


func GetCount(str string) (count int) {  // (count int) will init var with initial int val as 0
  for _, c := range str {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
      count++
    }
  }
  return count
}


func GetCount(strn string) int {
	count := 0
	vowels := []string{"a", "e", "i", "o", "u"}  //  create dynamically-sized slice
	for _, vowel := range vowels {
		count += strings.Count(strn, vowel)  // str.count(val)
	}
	return count
}

import "regexp"

func GetCount(str string) (count int) {
	r := regexp.MustCompile("[aeiou]")  // re_pattern = '[aeiou]'
	vowels := r.FindAllString(str, -1)  // re.find_all(re_pattern, string)
	return len(vowels)
}



// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 
// - - - - - - - - - - Shortest Word - - - - - - - - - - - - - 
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 

// Find shortest word in string of words

import (
  "strings"
)

func FindShort(s string) (min int) {
  words := strings.Split(s, " ")
  for _, word := range words {
    len_word := len(word)
    if len_word < min || min == 0 {
      min = len_word
    }
  }
  return 
}


import (
  "strings"
)

func FindShort(s string) (min int) {
  words := strings.Split(s, " ")  // cound also use strings.Fields(s)  split a str into substrings based on white space characters (spaces, tabs, newlines, carriage returns, form feeds, and vertical tabs)
  for _, word := range words {
    len_word := len(word)
    if len_word < min || min == 0 {
      min = len_word
    }
  }
  return 
}

