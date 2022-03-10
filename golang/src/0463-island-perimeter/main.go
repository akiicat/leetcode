package main

// T: O(m*n)
// M: O(1)
// -- start --

// T: O(m*n)
// M: O(1)
func islandPerimeter(grid [][]int) int {
  sum := 0
  h, w := len(grid), len(grid[0])
  for i, row := range grid {
    for j, col := range row {
      if col == 0 {
        continue
      }

      // up
      if i == 0 || grid[i-1][j] == 0 {
        sum++
      }

      // down
      if i == h-1 || grid[i+1][j] == 0 {
        sum++
      }

      // left
      if j == 0 || grid[i][j-1] == 0 {
        sum++
      }

      // right
      if j == w-1 || grid[i][j+1] == 0 {
        sum++
      }

    }
  }
  return sum
}

// T: O(m*n)
// M: O(1)
func islandPerimeterRemove(grid [][]int) int {
  c, sum := 0, 0
  for i, row := range grid {
    for j, col := range row {
      if col == 1 {
        c++
        sum += NearByOne(grid, i, j)
      }
    }
  }
  return c << 2 - sum
}

func NearByOne(grid [][]int, i, j int) int {
  rtn := 0

  x := len(grid)
  y := len(grid[0])

  if i > 0 {
    rtn += grid[i-1][j]
  }

  if i < x-1 {
    rtn += grid[i+1][j]
  }

  if j > 0 {
    rtn += grid[i][j-1]
  }

  if j < y-1 {
    rtn += grid[i][j+1]
  }

  return rtn
}

// -- end --

