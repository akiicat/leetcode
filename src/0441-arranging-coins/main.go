package main
import "math"

// T: O(log(n))
// M: O(1)
// -- start --

// Binary Search
// T: O(log(n))
// M: O(1)
func arrangeCoins(n int) int {
  if n == 1 {
    return 1
  }

  // low, high
  l, h := 1, n
  for l < h {
    m := (l + h) / 2
    if (m * (m + 1) >> 1) > n {
      h = m
    } else {
      l = m + 1
    }
  }

  return l - 1
}

// Bit Binary Search
// T: O(log(n))
// M: O(1)
func arrangeCoinsBit(n int) int {
  s := 0
  b := 0x8000
  for b != 0 {
    s = s ^ b    // add b
    if (s * (s + 1) >> 1) > n {   // test if s is too large after s add b
      s = s ^ b  // remove b if s is too large
    }
    b = b >> 1   // next b (half)
  }
  return s
}

// T: O(log(n)) log(n) for sqrt
// M: O(1)
func arrangeCoinsFormula(n int) int {
  return (int(math.Sqrt(8*float64(n) + 1)) - 1) >> 1
}

// -- end --

