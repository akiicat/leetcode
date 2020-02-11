package main

// leetcode 83.
// T: O(n)
// M: O(1)
// -- start --

func removeDuplicates(nums []int) int {
  cur := 0

  for _, num := range nums {
    if nums[cur] != num {
      cur++
      nums[cur] = num
    }
  }

  return cur + 1
}

// -- end --

