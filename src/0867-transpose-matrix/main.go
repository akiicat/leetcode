package main

// T: O(n * m)
// M: O(1)
// -- start --

func transpose(A [][]int) [][]int {
  n, m := len(A), len(A[0])

  T := make([][]int, m)
  for i := range T {
    T[i] = make([]int, n)
  }

  for i, row := range A {
    for j, v := range row {
      T[j][i] = v
    }
  }

  return T
}

// -- end --

