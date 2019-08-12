package main
import "fmt"

func main() {
  i, o := 5, [][]int{[]int{1}, []int{1,1}, []int{1,2,1}, []int{1,3,3,1}, []int{1,4,6,4,1}}
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %v\n", generate(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = 6, [][]int{[]int{1}, []int{1,1}, []int{1,2,1}, []int{1,3,3,1}, []int{1,4,6,4,1}, []int{1,5,10,10,5,1}}
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %v\n", generate(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = 1, [][]int{[]int{1}}
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %v\n", generate(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = 0, [][]int{}
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %v\n", generate(i))
  fmt.Printf("Expect: %v\n", o)
}

// https://leetcode.com/articles/pascals-triangle/
// T: O(n ** 2) n for numRows
// M: O(n ** 2)
// -- start --

// Dynamic Programming
func generate(numRows int) [][]int {
  if numRows == 0 {
    return [][]int{}
  }

  tri := make([][]int, numRows)
  tri[0] = make([]int, 1)
  tri[0][0] = 1

  for r := 1; r < numRows; r++ {
    tri[r] = make([]int, r + 1)
    tri[r][0] = 1
    tri[r][r] = 1

    for c := 1; c < r; c++ {
      tri[r][c] = tri[r-1][c-1] + tri[r-1][c]
    }
  }

  return tri
}

// -- end --

