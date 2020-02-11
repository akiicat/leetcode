package main
import "math/bits"

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

