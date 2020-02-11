package main

// T: O(n)
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func rotate(nums []int, k int)  {
  l := len(nums)
  k %= l
  copy(nums, append(nums[l-k:], nums[:l-k]...))
}

// -- end --

