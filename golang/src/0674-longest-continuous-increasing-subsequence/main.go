package main

// T: O(n)
// M: O(1)
// -- start --

func findLengthOfLCIS(nums []int) int {
  if len(nums) == 0 {
    return 0
  }

  max, cur := 1, 1

  for i := 1; i < len(nums); i++ {
    if nums[i-1] >= nums[i] {
      max = Max(max, cur)
      cur = 0
    }
    cur++
  }

  return Max(max, cur)
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

