package main
import "fmt"

func main() {
  i, k, o := []int{1}, 0, 0
  fmt.Printf("Input:  %v %d\n", i, k)
  fmt.Printf("Output: %d\n", smallestRangeI(i, k))
  fmt.Printf("Expect: %d\n", o)

  i, k, o = []int{0,10}, 0, 10
  fmt.Printf("Input:  %v %d\n", i, k)
  fmt.Printf("Output: %d\n", smallestRangeI(i, k))
  fmt.Printf("Expect: %d\n", o)

  i, k, o = []int{0,10}, 2, 6
  fmt.Printf("Input:  %v %d\n", i, k)
  fmt.Printf("Output: %d\n", smallestRangeI(i, k))
  fmt.Printf("Expect: %d\n", o)

  i, k, o = []int{0,10}, 10, 0
  fmt.Printf("Input:  %v %d\n", i, k)
  fmt.Printf("Output: %d\n", smallestRangeI(i, k))
  fmt.Printf("Expect: %d\n", o)

  i, k, o = []int{1,3,6}, 3, 0
  fmt.Printf("Input:  %v %d\n", i, k)
  fmt.Printf("Output: %d\n", smallestRangeI(i, k))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func smallestRangeI(A []int, K int) int {
  max := 0
  min := 1<<31

  for _, v := range A {
    max = Max(max, v)
    min = Min(min, v)
  }

  return Max(0, (max - min) - (2 * K))
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func Min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

