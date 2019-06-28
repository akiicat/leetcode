package main
import "fmt"
// import "math/bits"

func main() {
  fmt.Printf("Input:  0\nOutput: %t\nExpect: true\n", isPalindrome(0))
  fmt.Printf("Input:  1\nOutput: %t\nExpect: true\n", isPalindrome(1))
  fmt.Printf("Input:  121\nOutput: %t\nExpect: true\n", isPalindrome(121))
  fmt.Printf("Input:  -121\nOutput: %t\nExpect: false\n", isPalindrome(-121))
  fmt.Printf("Input:  313\nOutput: %t\nExpect: true\n", isPalindrome(313))
  fmt.Printf("Input:  1122\nOutput: %t\nExpect: false\n", isPalindrome(1122))
}

// T: O(N)
// M: O(1)
// -- start --

func isPalindrome(x int) bool {
  if x < 0 {
    return false
  }

  y := 0
  tmp := x
  for tmp != 0 {
    y = 10 * y + tmp % 10
    tmp /= 10
  }

  return x == y
}

// -- end --

