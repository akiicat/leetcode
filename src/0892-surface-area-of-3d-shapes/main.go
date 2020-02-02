package main
import "fmt"

func main() {
  i, o := [][]int{
    []int{2},
  }, 10
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", surfaceArea(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = [][]int{
    []int{1,2},
    []int{3,4},
  }, 34
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", surfaceArea(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = [][]int{
    []int{1,0},
    []int{0,2},
  }, 16
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", surfaceArea(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = [][]int{
    []int{1,1,1},
    []int{1,0,1},
    []int{1,1,1},
  }, 32
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", surfaceArea(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = [][]int{
    []int{2,2,2},
    []int{2,1,2},
    []int{2,2,2},
  }, 46
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", surfaceArea(i))
  fmt.Printf("Expect: %d\n", o)
}

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

