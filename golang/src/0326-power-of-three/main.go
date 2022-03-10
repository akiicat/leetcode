package main
import "math"

// leetcode 50. pow x n
// leetcode 231. power of two
// leetcode 342. power of four
// T: O(1)
// M: O(1)
// -- start --

func isPowerOfThree(n int) bool {
  // int(math.Pow(3, 20)) == 3486784401
  // int(math.Pow(2, 31)) == 2147483648
  return n > 0 && 3486784401 % n == 0
}

func isPowerOfThreeLog(n int) bool {
  if n == 0 {
    return false
  }
  res := math.Log(float64(n)) / math.Log(3)
  return (res - float64(int(res))) < 1E-14
}

func isPowerOfThreeForLoop(n int) bool {
  if n <= 0 {
    return false
  }

  for n != 1 {
    if n % 3 != 0 {
      return false
    }
    n /= 3
  }

  return true
}

// -- end --

