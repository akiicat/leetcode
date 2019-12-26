package main
import "fmt"

func main() {
  i, n, o := []int{1,3,5,6}, 5, 2
  fmt.Printf("Input:  %v, %d\n", i, n)
  fmt.Printf("Output: %d\n", searchInsert(i, n))
  fmt.Printf("Expect: %d\n", o)

  i, n, o = []int{1,3,5,6}, 2, 1
  fmt.Printf("Input:  %v, %d\n", i, n)
  fmt.Printf("Output: %d\n", searchInsert(i, n))
  fmt.Printf("Expect: %d\n", o)

  i, n, o = []int{1,3,5,6}, 7, 4
  fmt.Printf("Input:  %v, %d\n", i, n)
  fmt.Printf("Output: %d\n", searchInsert(i, n))
  fmt.Printf("Expect: %d\n", o)

  i, n, o = []int{1,3,5,6}, 0, 0
  fmt.Printf("Input:  %v, %d\n", i, n)
  fmt.Printf("Output: %d\n", searchInsert(i, n))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func searchInsert(nums []int, target int) int {
  for i, num := range nums {
    if num >= target {
      return i
    }
  }

  return len(nums)
}

// -- end --

