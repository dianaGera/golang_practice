package challange

func generate_arrey_between(a, b int) (arr []int) {
	// Complete the function that takes two integers 
	// (a, b, where a < b) and return an array of all 
	// integers between the input parameters, including them.
	for (a <= b) {
		arr = append(arr, a)
		a++
	}
	return
}


func CountPositivesSumNegatives(numbers []int) []int {
	// Return an array, where the first element is the count of positives 
	// numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.
	// If the input is an empty array or is null, return an empty array.
	// For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].
	res := make([]int, 2, 2)
	for _, v := range numbers {
	  if v < 0 {
		res[1] += v
	  } else if v > 0 {
		  res[0] += 1
	  }
	}
	return res
  }

