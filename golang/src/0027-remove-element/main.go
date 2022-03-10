package main

// T: O(n)
// M: O(1)
// -- start --

func removeElement(nums []int, val int) int {
  cur := 0

  for i, num := range nums {
    if num != val {
      nums[cur] = nums[i]
      cur++
    }
  }

  return cur
}

// -- end --

