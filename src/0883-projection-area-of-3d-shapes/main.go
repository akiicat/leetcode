package main

// T: O(n ** 2)
// M: O(1)
// -- start --

func projectionArea(grid [][]int) int {
  sum := 0

  for i, row := range grid {
    yz, xz := 0, 0
    for j, v := range row {
      if v != 0 {
        sum++
      }

      if v > yz {
        yz = v
      }

      if grid[j][i] > xz {
        xz = grid[j][i]
      }
    }
    sum += yz + xz
  }

  return sum
}

// -- end --

