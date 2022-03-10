package main
import "sort"

// T: O(n * log(n))
// M: O(1)
// -- start --

func minSubsequence(nums []int) []int {
  sort.Ints(nums)

  sum := 0
  for _, v := range nums {
    sum += v
  }

  rtn, s := []int{}, 0
  for i := len(nums)-1; s <= sum/2; i-- {
    s += nums[i]
    rtn = append(rtn, nums[i])
  }

  return rtn
}

// -- end --

