package main

// T: O(n)
// M: O(n)
// -- start --

func findLHS(nums []int) int {
  m := make(map[int]int)

  for _, v := range nums {
    m[v]++
  }

  max := 0
  for k, v := range m {
    next, ok := m[k+1]
    if ok && v + next > max {
      max = v + next
    }
  }

  return max
}

// -- end --

