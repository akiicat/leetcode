package main

// T: O(n + m)
// M: O(m)
// -- start --

func smallerNumbersThanCurrent(nums []int) []int {
  m := [101]int{}

  for _, v := range nums {
    m[v]++
  }

  for i := 1; i < 101; i++ {
    m[i] += m[i-1]
  }

  for i, v := range nums {
    if v != 0 {
      nums[i] = m[v-1]
    }
  }

  return nums
}

// -- end --

