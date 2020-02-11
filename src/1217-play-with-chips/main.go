package main

// T: O(n)
// M: O(1)
// -- start --

func minCostToMoveChips(chips []int) int {
  m := []int{0,0}
  for _, v := range chips {
    m[v & 0x1]++
  }
  return Min(m[0], m[1])
}

func Min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

