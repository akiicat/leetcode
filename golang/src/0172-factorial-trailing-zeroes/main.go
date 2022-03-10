package main

// T: O(log(n))
// M: O(1)
// -- start --

func trailingZeroes(n int) int {
  sum := 0

  for {
    n /= 5
    sum += n
    if n == 0 {
      break
    }
  }

  return sum
}

// -- end --

