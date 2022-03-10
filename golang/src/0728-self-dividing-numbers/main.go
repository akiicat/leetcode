package main

// T: O(n) n is the number of integers in the range
// M: O(1)
// -- start --

func selfDividingNumbers(left int, right int) []int {
  m := []int{}

  for i := left; i <= right; i++ {
    if Div(i) {
      m = append(m, i)
    }
  }

  return m
}

func Div(i int) bool {
  if i < 10 {
    return true
  }

  n := i

  for n != 0 {
    mod := n % 10
    if mod == 0 || i % mod != 0 {
      return false
    }
    n /= 10
  }

  return true
}

// -- end --

