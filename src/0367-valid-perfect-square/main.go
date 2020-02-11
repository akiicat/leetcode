package main
import "math"

// leetcode 69.
// T: O(log(n)) // the time complexity of Sqrt function if log(n)
// M: O(1)
// -- start --

// T: O(sqrt(n))
// M: O(1)
func isPerfectSquare(num int) bool {
	i := 1
	for num > 0 {
		num -= i
		i += 2
	}
	return num == 0
}

// T: O(log(n))
// M: O(1)
func isPerfectSquareMath(num int) bool {
  sqrt := math.Sqrt(float64(num))
  return sqrt == float64(int(sqrt))
}

// -- end --

