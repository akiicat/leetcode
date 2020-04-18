package main

// T: O(n)
// M: O(1)
// -- start --

func minStartValue(nums []int) int {
  min, sum := -1, -1
  for _, v := range nums {
    sum += v
    if sum < min {
      min = sum
    }
  }
  return -min
}

// -- end --

