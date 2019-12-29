package main
import "fmt"

func main() {
  i, o := []int{1,7,3,6,5,6}, 3
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", pivotIndex(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,2,3}, -1
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", pivotIndex(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func pivotIndex(nums []int) int {
  l, r := 0, 0

  for _, v := range nums {
    r += v
  }

  for i, v := range nums {
    r -= v
    if l == r {
      return i
    }
    l += v
  }

  return -1
}

// -- end --

