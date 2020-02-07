package main
import "fmt"

func main() {
  i, o := []int{3,1,2,4}, []int{2,4,3,1}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", sortArrayByParity(i))
  fmt.Printf("Expect: %v\n", o)
}

// leetcode 922.
// T: O(n)
// M: O(1)
// -- start --

func sortArrayByParity(A []int) []int {
  i := 0
  for j, v := range A {
    if v & 0x1 == 0 {
      A[i], A[j] = v, A[i]
      i++
    }
  }

  return A
}

// -- end --

