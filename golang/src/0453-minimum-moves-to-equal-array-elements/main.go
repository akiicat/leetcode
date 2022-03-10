package main

// T: O(n)
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func minMoves(nums []int) int {
  sum := 0
  min := 1<<31
  for _, num := range nums {
    sum += num
    if num < min {
      min = num
    }
  }

  return sum - len(nums) * min
}

// T: O(n)
// M: O(1)
func minMovesDiff(nums []int) int {
  min := 1<<31
  for _, num := range nums {
    if num < min {
      min = num
    }
  }

  sum := 0
  for _, num := range nums {
    sum += num - min
  }

  return sum
}

// -- end --

