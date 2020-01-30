package main
import "fmt"

func main() {
  i, o := [][]int{
    []int{1,2,3},
    []int{4,5,6},
    []int{7,8,9},
  }, [][]int{
    []int{1,4,7},
    []int{2,5,8},
    []int{3,6,9},
  }
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", transpose(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = [][]int{
    []int{1,2,3},
    []int{4,5,6},
  }, [][]int{
    []int{1,4},
    []int{2,5},
    []int{3,6},
  }
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", transpose(i))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n * m)
// M: O(1)
// -- start --

func transpose(A [][]int) [][]int {
  n, m := len(A), len(A[0])

  T := make([][]int, m)
  for i := range T {
    T[i] = make([]int, n)
  }

  for i, row := range A {
    for j, v := range row {
      T[j][i] = v
    }
  }

  return T
}

// -- end --

