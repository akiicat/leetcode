package main

// T: O(n)
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func rob(nums []int) int {
  last, result := 0, 0
  for _, v := range nums {
    tmp := result
    result = Max(last + v, result)
    last = tmp
  }
  return result
}


// T: O(n ** 2)
// M: O(n ** 2)
func robRecursive(nums []int) int {
  if len(nums) == 0 {
    return 0
  }

  if len(nums) == 1 {
    return nums[0]
  }

  // take: nums[0] + rob(nums[2:])
  // not take: rob(nums[1:])
  return Max(nums[0] + rob(nums[2:]), rob(nums[1:]))
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

