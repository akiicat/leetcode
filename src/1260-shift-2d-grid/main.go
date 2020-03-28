package main

// T: O(n)
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func shiftGrid(grid [][]int, k int) [][]int {
  n, m := len(grid), len(grid[0])

  rtn := make([][]int, n)
  for i := 0; i < n; i++ {
    rtn[i] = make([]int, m)
  }

  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      y := k / m + i
      x := k % m + j
      if x >= m {
        y++
      }
      rtn[y%n][x%m] = grid[i][j]
    }
  }

  return rtn
}

// T: O(n)
// M: O(n)
func shiftGridArray(grid [][]int, k int) [][]int {
  n, m := len(grid), len(grid[0])

  q := []int{}
  for _, v := range grid {
    q = append(q, v...)
  }

  k = len(q) - k % len(q)
  q = append(q[k:], q[:k]...)

  grid = [][]int{}
  for i := 0; i < n; i++ {
    grid, q = append(grid, q[:m]), q[m:]
  }

  return grid
}

// -- end --

