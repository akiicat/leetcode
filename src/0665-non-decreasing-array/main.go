package main

// T: O(n)
// M: O(1)
// -- start --

func checkPossibility(nums []int) bool {
  if len(nums) < 3 {
    return true
  }

  c := false
  for i := 0; i < len(nums) - 1; i++ {
    if nums[i] <= nums[i+1] {
      continue
    }

    if (i != 0 && nums[i-1] > nums[i+1]) && (i + 2 != len(nums) && nums[i] > nums[i+2]) {
      return false
    }

    if c {
      return false
    }

    c = true
  }

  return true
}

// -- end --

