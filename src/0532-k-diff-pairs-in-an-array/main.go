package main

// T: O(n)
// M: O(n)
// -- start --

func findPairs(nums []int, k int) int {
  if k < 0 {
    return 0
  }

  m := make(map[int]int)
  for _, num := range nums {
    m[num]++
  }

  c := 0

  for _, num := range nums {
    if k == 0 {
      if m[num] >= 2{
        c++
        m[num] = 0
      }
    } else if m[num+k] > 0 {
      c++
      m[num+k] = 0
    }
  }

  return c
}

// -- end --

