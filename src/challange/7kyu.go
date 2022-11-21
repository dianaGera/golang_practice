package challange

import (
	"math"
)


// Insertion sort
func SortNumbers(numbers []int) []int {
	// solution(c(1, 2, 3, 10, 5)) # should return c(1, 2, 3, 5, 10)
	// solution(NULL)              # should return NULL
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
		return numbers
	} else {
		return []int{} 
	}
}



func FindNextSquare(sq int64) int64 {
	// 121 --> 144
	// 625 --> 676
	// 114 --> -1 since 114 is not a perfect square
	i := math.Sqrt(float64(sq))
	d := math.Floor(i)
	remainder := i - d
	if remainder == 0 {
	  d++
	  res := int64(math.Pow(d, 2))
	  return res
	}
	return -1
  }