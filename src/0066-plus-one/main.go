package main
import "fmt"

func main() {
  i, o := []int{1,2,3}, []int{1,2,4}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\nExpect: %v\n", plusOne(i), o)

  i, o = []int{4,3,2,1}, []int{4,3,2,2}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\nExpect: %v\n", plusOne(i), o)
}

// T: O(N)
// M: O(1)
// -- start --

func plusOne(digits []int) []int {
  carry := 1
  for i := len(digits) - 1; i >= 0; i-- {
    digits[i] += carry

    // digits[i], carry = sum % 10, sum / 10
    if digits[i] >= 10 {
      digits[i] -= 10
    } else {
      carry = 0
    }

  }

  if carry == 1 {
    return append([]int{carry}, digits...)
  }

  return digits
}

// -- end --

