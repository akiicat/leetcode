package main
import "fmt"
import "math/bits"

func main() {
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 0, hasAlternatingBits(0), true)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 1, hasAlternatingBits(1), true)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 2, hasAlternatingBits(2), true)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 3, hasAlternatingBits(3), false)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 4, hasAlternatingBits(4), false)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 5, hasAlternatingBits(5), true)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 1431655765, hasAlternatingBits(1431655765), true)
}

// T: O(1)
// M: O(1)
// -- start --

func hasAlternatingBits(n int) bool {
  return bits.OnesCount64(uint64(n ^ (n << 1 | ((n ^ 0x1) & 0x1))) + 1) == 1
}

// https://leetcode.com/problems/binary-number-with-alternating-bits/discuss/108427/Oneliners-(C%2B%2B-Java-Ruby-Python)
//     000101010
//   ^ 000001010
//   = 000100000
func hasAlternatingBitsCancelBits(n int) bool {
  n = n ^ n >> 2
  return n & (n - 1) == 0
}

//     000101010
//   ^ 000010101
//   = 000111111
func hasAlternatingBitsCompleteBits(n int) bool {
  n = n ^ n >> 1
  return n & (n + 1) == 0
}

// -- end --

