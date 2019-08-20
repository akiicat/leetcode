package main
import "fmt"

func main() {
  i, o := 3, []int{1,3,3,1}
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %v\n", getRow(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = 0, []int{1}
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %v\n", getRow(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = 4, []int{1,4,6,4,1}
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %v\n", getRow(i))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n ** 2)
// M: O(n)
// -- start --

// Dynamic Programming
// T: O(n ** 2)
// M: O(n)
func getRow(rowIndex int) []int {
  tri := make([]int, rowIndex + 1)
  tri[0] = 1

  for r := 1; r < rowIndex + 1; r++ {
    for c := r; c > 0; c-- {
      tri[c] = tri[c] + tri[c-1]
    }
  }

  return tri
}

// -- end --

