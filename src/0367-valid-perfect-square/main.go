package main
import "fmt"
import "math"

func main() {
  i, o := 16, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", isPerfectSquare(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 14, false
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", isPerfectSquare(i))
  fmt.Printf("Expect: %t\n", o)
}

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

