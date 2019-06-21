package main
import "fmt"

func main() {
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 0, isPowerOfTwo(0), false)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 1, isPowerOfTwo(1), true)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 2, isPowerOfTwo(2), true)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 3, isPowerOfTwo(3), false)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 4, isPowerOfTwo(4), true)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 16, isPowerOfTwo(16), true)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 218, isPowerOfTwo(218), false)
}

// T: O(1)
// M: O(1)
// -- start --

func isPowerOfTwo(n int) bool {
  return n > 0 && n & (n - 1) == 0
}

// -- end --

