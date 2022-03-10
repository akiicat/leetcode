package main

import "math/rand"

// T: O(1)
// M: O(1)
// -- start --

func getNoZeroIntegers(n int) []int {
  a, b := 1, n-1

  for Zero(a) || Zero(b) {
    a = 1 + rand.Intn(n-2)
    b = n - a
  }

  return []int{a,b}
}

func Zero(n int) bool {
  for n != 0 {
    if n % 10 == 0 {
      return true
    }
    n /= 10
  }
  return false
}

// -- end --

