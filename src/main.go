package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 10, 50, 5}
	if cap(numbers) > 0 {
		
		for i := range numbers {
		  y := i
		  for y > 0 && numbers[y] < numbers[y-1] {
			temp := numbers[y]
			numbers[y] = numbers[y-1]
			numbers[y-1] = temp
			y -= 1
		  }
		}
		fmt.Println(numbers) 
	  }
	
  
}
