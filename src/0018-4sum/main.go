package main
import "sort"

// T: O(n ^ 3)
// M: O(1)
// -- start --

func fourSum(nums []int, target int) [][]int {
  sort.Ints(nums)

  res := [][]int{}
  for i1 := 0; i1 < len(nums)-3; i1++ {
    if i1 > 0 && nums[i1] == nums[i1-1] {
      continue
    }

    for i2 := i1+1; i2 < len(nums)-2; i2++ {
      if i2 > i1+1 && nums[i2] == nums[i2-1] {
        continue
      }

      l, r := i2+1, len(nums)-1
      for l < r {
        sum := nums[i1] + nums[i2] + nums[l] + nums[r]

        if sum == target {
          res = append(res, []int{nums[i1], nums[i2], nums[l], nums[r]})
          l, r = l+1, r-1

          for l < r && nums[l] == nums[l-1] {
            l++
          }

          for l < r && nums[r] == nums[r+1] {
            r--
          }
        } else if sum > target {
          r--
        } else {
          l++
        }

      }
    }
  }

  return res
}

// -- end --

