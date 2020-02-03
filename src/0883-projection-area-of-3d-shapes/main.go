package main
import "fmt"

func main() {
  i, o := [][]int{
    []int{2},
  }, 5
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", projectionArea(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = [][]int{
    []int{1,2},
    []int{3,4},
  }, 17
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", projectionArea(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = [][]int{
    []int{1,0},
    []int{0,2},
  }, 8
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", projectionArea(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = [][]int{
    []int{1,1,1},
    []int{1,0,1},
    []int{1,1,1},
  }, 14
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", projectionArea(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = [][]int{
    []int{2,2,2},
    []int{2,1,2},
    []int{2,2,2},
  }, 21
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", projectionArea(i))
  fmt.Printf("Expect: %d\n", o)
}

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

