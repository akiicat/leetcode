package main

// leetcode 414. 506.
// T: O(n)
// M: O(1)
// -- start --

func maximumProduct(nums []int) int {
  if len(nums) == 3 {
    return nums[0] * nums[1] * nums[2]
  }

  max := [3]int{}
  min := [2]int{}

  for _, v := range nums {
    if v > 0 {
      if v > max[0] {
        max[0], max[1], max[2] = v, max[0], max[1]
      } else if v > max[1] {
        max[1], max[2] = v, max[1]
      } else if v > max[2] {
        max[2] = v
      }
    } else {
      if v < min[0] {
        min[0], min[1] = v, min[0]
      } else if v < min[1] {
        min[1] = v
      }
    }
  }

  return Max(max[0] * max[1] * max[2], max[0] * min[0] * min[1])
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

