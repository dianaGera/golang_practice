// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 
// - - - - - - - - - - Persistent Bugger - - - - - - - - - - - - - - 
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 

// 39 --> 3 (because 3*9 = 27, 2*7 = 14, 1*4 = 4 and 4 has only one digit)
// 999 --> 4 (because 9*9*9 = 729, 7*2*9 = 126, 1*2*6 = 12, and finally 1*2 = 2)
// 4 --> 0 (because 4 is already a one-digit number)
// check loading time


func Persistence(n int) int {
  steps := 0
	for n >= 10 {
    m := 1
    for n > 0 {
      m *= n % 10
      n /= 10
    }
    n = m
    steps++
  }
  return steps
}


// - - - - - - - - - recursion - - - - - - - - - - -

import "strconv"

func Persistence(n int) int {
  return Calc(n, 0)
}
func Calc(n int, i int) int {
  if (n < 10) {
    return i
  }
  nStr := strconv.Itoa(n)  // int(str)
  subResult := 1
  for _, char := range nStr {
    num, err := strconv.Atoi(string(char))  // str(int)
    if err != nil {
      panic(err)
    }
    subResult = num * subResult
  }
  return Calc(subResult, i + 1)
}
 
