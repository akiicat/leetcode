package main

// T: O(n)
// M: O(1)
// -- start --

func decompressRLElist(nums []int) []int {
  m := []int{}

  for i := 0; i < len(nums); i += 2 {
    for f := 0; f < nums[i]; f++ {
      m = append(m, nums[i+1])
    }
  }

  return m
}

// -- end --

