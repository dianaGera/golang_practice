package challange


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