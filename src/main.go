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
	var mapping map[rune]int = map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L':50, 'C':100, 'D':500, 'M':1000}

	fmt.Println(mapping['I'])
}
