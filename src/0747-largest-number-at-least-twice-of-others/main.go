package main

// T: O(n)
// M: O(1)
// -- start --

func dominantIndex(nums []int) int {
  if len(nums) == 1 {
    return 0
  }

  idx, m1, m2 := 0, 0, 0

  for i, v := range nums {
    if v > m1 {
      idx, m1, m2 = i, v, m1
    } else if v > m2 {
      m2 = v
    }
  }

  if m1 >= m2 << 1 {
    return idx
  }

  return -1
}

// -- end --

