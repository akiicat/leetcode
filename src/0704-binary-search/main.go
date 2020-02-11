package main

// T: O(log(n))
// M: O(1)
// -- start --

// T: O(log(n))
// M: O(1)
func search(nums []int, target int) int {
  l, m, r := 0, 0, len(nums) - 1

  for l < r {
    m = (l + r) / 2
    if nums[m] == target {
      return m
    }
    if nums[m] > target {
      r = m - 1
    } else {
      l = m + 1
    }
  }

  return -1
}

// T: O(log(n))
// M: O(log(n))
func searchRecursive(nums []int, target int) int {
  return R(nums, target, 0, len(nums) - 1)
}

func R(nums []int, target, l, r int) int {
  if l >= r {
    return -1
  }

  m := (l + r) / 2

  if nums[m] > target {
    return R(nums, target, l, m - 1)
  }

  if nums[m] < target {
    return R(nums, target, m + 1, r)
  }

  return m
}

// -- end --

