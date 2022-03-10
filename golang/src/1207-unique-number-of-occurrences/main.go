package main

// T: O(n)
// M: O(n)
// -- start --

func uniqueOccurrences(arr []int) bool {
  m := make(map[int]int)

  for _, v := range arr {
    m[v]++
  }

  n := make(map[int]bool)
  for _, v := range m {
    _, ok := n[v]
    if ok {
      return false
    }
    n[v] = true
  }

  return true
}

// -- end --

