package main

// T: O(n * m) n, n are the number of rows and columns
// M: O(1)
// -- start --

func numMagicSquaresInside(grid [][]int) int {
  sum := 0
  for i := 0; i < len(grid)-2; i++ {
    for j := 0; j < len(grid[i])-2; j++ {
      if isMagic(grid, i, j) {
        sum++
      }
    }
  }
  return sum
}

func isMagic(grid [][]int, x, y int) bool {
  m := 0
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      m ^= grid[x+i][y+j] ^ (3 * i + j + 1)
    }
  }

  if m != 0 {
    return false
  }

  a := [8]int{}

  a[0] = grid[x  ][y] + grid[x  ][y+1] + grid[x  ][y+2]
  a[1] = grid[x+1][y] + grid[x+1][y+1] + grid[x+1][y+2]
  a[2] = grid[x+2][y] + grid[x+2][y+1] + grid[x+2][y+2]

  a[3] = grid[x][y  ] + grid[x+1][y  ] + grid[x+2][y  ]
  a[4] = grid[x][y+1] + grid[x+1][y+1] + grid[x+2][y+1]
  a[5] = grid[x][y+2] + grid[x+1][y+2] + grid[x+2][y+2]

  a[6] = grid[x  ][y] + grid[x+1][y+1] + grid[x+2][y+2]
  a[7] = grid[x+2][y] + grid[x+1][y+1] + grid[x  ][y+2]

  for _, v := range a[1:] {
    if v != a[0] {
      return false
    }
  }

  return true
}

// -- end --

