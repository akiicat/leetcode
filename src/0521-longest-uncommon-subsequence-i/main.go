package main

// T: O(min(a, b))
// M: O(1)
// -- start --

func findLUSlength(a string, b string) int {
  if a == b {
    return -1
  }
  return Max(len(a), len(b))
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

