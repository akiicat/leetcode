package main

// leetcode 231. power of two
// leetcode 326. power of three
// leetcode 342. power of four
// T: O(1) log(n), n for 32 bits
// M: O(1)
// -- start --

// X^27
// = X * X ^26
// = X * (X^2)^13
// = X * (X^2) * (X^2)^12
// = X * (X^2) * (X^4)^6
// = X * (X^2) * (X^8)^3
// = X * (X^2) * (X^8) * (X^8)^2
// = X * (X^2) * (X^8) * (X^16)
// T: O(1) log(n), n for 32 bits
// M: O(1)
func myPow(x float64, n int) float64 {
  if n == 0 {
    return 1.0
  }

  r := false
  if n < 0 {
    n, r = -n, true
  }

  res, base := 1.0, x
  for n > 1 {
    if n & 0x1 == 0 {
      n, base = n >> 1, base * base
    } else {
      n, res = n - 1, res * base
    }
  }

  res *= base

  if r {
    return 1 / res
  }

  return res
}

// T: O(n)
// M: O(1)
func myPowLoop(x float64, n int) float64 {
  r := false
  if n < 0 {
    n, r = -n, true
  }

  rtn := 1.0
  for i := 0; i < n; i++ {
    rtn *= x
  }

  if r {
    return 1 / rtn
  }

  return rtn
}

// -- end --

