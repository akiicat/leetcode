package main

// T: O(n * m) n, m are the number of rows and columns in matrix
// M: O(1)
// -- start --

func isToeplitzMatrix(matrix [][]int) bool {
  m, n := len(matrix), len(matrix[0])

  for y := 0; y < m - 1; y++ {
    for x := 0; x < n - 1; x++ {
      if matrix[y][x] != matrix[y+1][x+1] {
        return false
      }
    }
  }

  return true
}

// -- end --

