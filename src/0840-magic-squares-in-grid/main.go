package main
import "fmt"

func main() {
  i, o := [][]int{
    []int{4,3,8,4},
    []int{9,5,1,9},
    []int{2,7,6,2},
  }, 1
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", numMagicSquaresInside(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = [][]int{
    []int{5,5,5},
    []int{5,5,5},
    []int{5,5,5},
  }, 0
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", numMagicSquaresInside(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = [][]int{
    []int{3,2,1,6},
    []int{5,9,6,8},
    []int{1,5,1,2},
    []int{3,7,3,4},
  }, 0
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", numMagicSquaresInside(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = [][]int{
    []int{1,3,7,8,8},
    []int{8,3,2,7,4},
    []int{3,8,4,0,9},
    []int{8,1,6,5,0},
    []int{7,2,1,8,6},
  }, 0
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", numMagicSquaresInside(i))
  fmt.Printf("Expect: %d\n", o)
}

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

