package main

// T: O(n + m)
// M: O(1)
// -- start --

// T: O(n + m)
// M: O(1)
func countNegatives(grid [][]int) int {
  m, n := len(grid), len(grid[0])

  sum, r, c := 0, 0, n - 1
  for (r < m) && (c >= 0) {
    if grid[r][c] < 0 {
      sum += m - r
      c--
    } else {
      r++
    }
  }

  return sum
}

// T: O(n*m)
// M: O(1)
func countNegativesRow(grid [][]int) int {
  m, n := len(grid), len(grid[0])

  sum := 0
  for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
      if grid[i][j] < 0 {
        sum += (n - j) * (m - i)
        n = j
      }
    }
  }

  return sum
}

// -- end --

