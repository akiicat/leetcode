package main

// leetcode 50. pow x n
// leetcode 326. power of three
// leetcode 342. power of four
// T: O(1)
// M: O(1)
// -- start --

func isPowerOfTwo(n int) bool {
  return n > 0 && n & (n - 1) == 0
}

// -- end --

