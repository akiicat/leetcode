package main

// T: O(n)
// M: O(1)
// -- start --

func sumZero(n int) []int {
  m := make([]int, n)

  u := 1
  for i := 0; i < len(m)-1; i += 2 {
    m[i], m[i+1], u = u, -u, u+1
  }

  return m
}

// -- end --

