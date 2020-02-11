package main

// T: O(n)
// M: O(1)
// -- start --

func searchInsert(nums []int, target int) int {
  for i, num := range nums {
    if num >= target {
      return i
    }
  }

  return len(nums)
}

// -- end --

