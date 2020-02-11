package main

// https://leetcode.com/articles/missing-number/
// T: O(n)
// M: O(1)
// -- start --

// Gauss' Formula
// T: O(n)
// M: O(1)
func missingNumber(nums []int) int {
  sum := 0
  for _, v := range nums {
    sum += v
  }
  size := len(nums)
  missing := (1 + size) * size / 2
  return missing - sum
}


// Bit Manipulation
// T: O(n)
// M: O(1)
func missingNumberBitManipulation(nums []int) int {
  missing := len(nums)
  for i, v := range nums {
    missing = missing ^ i ^ v
  }
  return missing
}

// -- end --
