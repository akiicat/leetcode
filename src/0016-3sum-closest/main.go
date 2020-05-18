package main
import "sort"

// T: O(n ^ 2)
// M: O(1)
// -- start --

func threeSumClosest(nums []int, target int) int {
  sort.Ints(nums)

  min, res := 1<<31, 0
  for i := 0; i < len(nums)-2; i++ {
    if i > 0 && nums[i] == nums[i-1] {
      continue
    }

    l, r := i+1, len(nums)-1
    for l < r {
      d := nums[i] + nums[l] + nums[r] - target

      if Abs(d) < min {
        min, res = Abs(d), nums[i] + nums[l] + nums[r]
      } else if d > 0 {
        r--
      } else {
        l++
      }
    }
  }
  return res
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func Abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}

// -- end --

