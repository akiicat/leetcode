package main

// T: O(n)
// M: O(1)
// -- start --

func findMaxAverage(nums []int, k int) float64 {
  if len(nums) < k {
    return 0
  }

  sum := 0
  i := 0
  for i < k {
    sum += nums[i]
    i++
  }

  max := sum
  for i < len(nums) {
    sum += nums[i] - nums[i-k]
    if sum > max {
      max = sum
    }
    i++
  }

  return float64(max) / float64(k)
}

// -- end --

