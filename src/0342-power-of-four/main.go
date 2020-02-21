package main
import "math"
import "strconv"

// leetcode 50. pow x n
// leetcode 231. power of two
// leetcode 326. power of three
// T: O(1)
// M: O(1)
// -- start --

func isPowerOfFour(num int) bool {
  if num <= 0 {
    return false
  }

  s := strconv.FormatInt(int64(num), 2)
  length := len(s)
  shiftValue, _ := strconv.Atoi(s[1:])

  return shiftValue == 0 && length % 2 == 1
}

func isPowerOfFourLog(num int) bool {
  if num <= 0 {
    return false
  }
  res := math.Log(float64(num)) / math.Log(4)
  return (res - float64(int(res))) < 1E-14
}

func isPowerOfFourForLoop(n int) bool {
  if n <= 0 {
    return false
  }

  for n != 1 {
    if n % 4 != 0 {
      return false
    }
    n >>= 2
  }

  return true
}

// -- end --

