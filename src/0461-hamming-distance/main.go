package main
import "math/bits"

// T: O(1)
// M: O(1)
// -- start --

func hammingDistance(x int, y int) int {
  return bits.OnesCount(uint(x ^ y))
}

// -- end --

