// Beta
// Given a 2d matrix of m columns and n rows.
// The matrix is filled with 0,1,2,3,.. , from left to right, top to down.
// where A-M-D and C-M-B are diagonal in a smaller 3x3 matrix.

// Tips:

// Both m and n are in range [3, 25) or [3, 100), vary among languages.
// Inputsum >= 0.
// A valid sum's M should be within the matrix, as well as its A, B, C, D.
// For any invalid input, return false.

// Expect(CrossSumCheck(3, 3, 5)).To(Equal(false))
// Expect(CrossSumCheck(4, 3, 25)).To(Equal(true))
// Expect(CrossSumCheck(4, 3, 26)).To(Equal(false))
// Expect(CrossSumCheck(3, 4, 20)).To(Equal(true))
// Expect(CrossSumCheck(3, 4, 25)).To(Equal(false))

package challange

import (
	"fmt"
)

var m = 3
var n = 4
var sum = 20

func find_sum_of_shape_x_in_matrix() bool {

	if n >= 3 {
		count := 0
		arr := make([][]int, n)
		for i := range arr {
			arr[i] = make([]int, m)
			for j := range arr[i] {
				arr[i][j] = count
				count += 1
			}
		}

		for i := range arr {

			fmt.Println(arr[i])
		}

		var sum_arr = 0
		var i = 0
		for i <= 2 {
			if i == 0 || i/2 == 1 {
				sum_arr += arr[i][0] + arr[i][2]
			} else {
				sum_arr += arr[i][i]
			}
			i++
		}

		if sum_arr == sum {
			return true
		} else {
			return false
		}

	} else {
		return false
	}

}
