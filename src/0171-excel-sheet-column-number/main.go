package main

// T: O(1)
// M: O(1)
// -- start --

func titleToNumber(s string) int {
  n := 0
  for _, v := range s {
    n = 26 * n + int(v) - 64
  }
  return n
}

// -- end --

