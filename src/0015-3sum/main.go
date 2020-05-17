package main
import "sort"

// T: O(n ^ 2)
// M: O(1)
// -- start --

func threeSum(nums []int) [][]int {
  sort.Ints(nums)

  res := [][]int{}
  for i := 0; i < len(nums)-2; i++ {
    if i > 0 && nums[i] == nums[i-1] {
      continue
    }

    l, r := i+1, len(nums)-1
    for l < r {
      if nums[i] + nums[l] + nums[r] == 0 {
        res = append(res, []int{nums[i], nums[l], nums[r]})
        l, r = l+1, r-1

        for l < r && nums[l] == nums[l-1] {
          l++
        }

        for l < r && nums[r] == nums[r+1] {
          r--
        }
      } else if nums[i] + nums[l] + nums[r] > 0 {
        r--
      } else {
        l++
      }
    }

  }

  return res
}

// -- end --

