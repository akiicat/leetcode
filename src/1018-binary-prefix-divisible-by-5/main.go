package main

// T: O(n)
// M: O(1)
// -- start --

func prefixesDivBy5(A []int) []bool {
  m := make([]bool, len(A))
  n := 0

  for i, v := range A {
    n = (n << 1 | v) % 5

    if n == 0 {
      m[i] = true
    }
  }

  return m
}

// -- end --

