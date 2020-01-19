package main
import "fmt"

func main() {
  i, o := []int{1,2,3}, 1
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", minCostToMoveChips(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{2,2,2,3,3}, 2
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", minCostToMoveChips(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,2,2,2,2}, 1
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", minCostToMoveChips(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func minCostToMoveChips(chips []int) int {
  m := []int{0,0}
  for _, v := range chips {
    m[v & 0x1]++
  }
  return Min(m[0], m[1])
}

func Min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

