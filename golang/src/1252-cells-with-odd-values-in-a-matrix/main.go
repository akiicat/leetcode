package main

// T: O(n * m)
// M: O(n + m)
// -- start --

func oddCells(n int, m int, indices [][]int) int {
  r := make([]int, n)
  c := make([]int, m)

  for _, v := range indices {
    r[v[0]]++
    c[v[1]]++
  }

  odds := 0
  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      if (r[i] + c[j]) & 0x1 == 1 {
        odds++
      }
    }
  }

  return odds
}

// -- end --

