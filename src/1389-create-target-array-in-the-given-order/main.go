package main

// T: O(n)
// M: O(1)
// -- start --

func createTargetArray(nums []int, index []int) []int {
  m := make([]int, len(nums))

  for k := 0; k < len(nums); k++ {
    n, i := nums[k], index[k]

    copy(m[i+1:], m[i:])
    m[i] = n
  }

  return m
}

// -- end --

