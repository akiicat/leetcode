package main
import "fmt"

func main() {
  i, o := [][]int{[]int{1,2,3,4},[]int{5,1,2,3},[]int{9,5,1,2}}, true
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", isToeplitzMatrix(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = [][]int{[]int{1,2},[]int{2,2}}, false
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", isToeplitzMatrix(i))
  fmt.Printf("Expect: %t\n", o)
}

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

