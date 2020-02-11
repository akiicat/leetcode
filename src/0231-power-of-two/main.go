package main

// T: O(1)
// M: O(1)
// -- start --

func isPowerOfTwo(n int) bool {
  return n > 0 && n & (n - 1) == 0
}

// -- end --

