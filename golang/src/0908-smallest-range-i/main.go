package main

// T: O(n)
// M: O(1)
// -- start --

func smallestRangeI(A []int, K int) int {
  max := 0
  min := 1<<31

  for _, v := range A {
    max = Max(max, v)
    min = Min(min, v)
  }

  return Max(0, (max - min) - (2 * K))
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func Min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

