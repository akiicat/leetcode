package main

// T: O(n)
// M: O(1)
// -- start --

func pivotIndex(nums []int) int {
  l, r := 0, 0

  for _, v := range nums {
    r += v
  }

  for i, v := range nums {
    r -= v
    if l == r {
      return i
    }
    l += v
  }

  return -1
}

// -- end --

