package main

// T: O(n ** 2)
// M: O(1)
// -- start --

func surfaceArea(grid [][]int) int {
  sum := 0

  for i, row := range grid {
    for j, v := range row {
      if v != 0 {
        sum += 4 * v + 2
      }

      if i != 0 {
        sum -= 2 * Min(v, grid[i-1][j])
      }

      if j != 0 {
        sum -= 2 * Min(v, grid[i][j-1])
      }
    }
  }

  return sum
}

func Min(a, b int) int {
  if a > b {
    return b
  }
  return a
}

// -- end --

