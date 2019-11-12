package main
import "fmt"

func main() {
  i, k, o := []int{3,1,4,1,5}, 2, 2
  fmt.Printf("Input:  %v %d\n", i, k)
  fmt.Printf("Output: %d\n", findPairs(i, k))
  fmt.Printf("Expect: %d\n", o)

  i, k, o = []int{1,2,3,4,5}, 1, 4
  fmt.Printf("Input:  %v %d\n", i, k)
  fmt.Printf("Output: %d\n", findPairs(i, k))
  fmt.Printf("Expect: %d\n", o)

  i, k, o = []int{1,3,1,5,4}, 0, 1
  fmt.Printf("Input:  %v %d\n", i, k)
  fmt.Printf("Output: %d\n", findPairs(i, k))
  fmt.Printf("Expect: %d\n", o)

  i, k, o = []int{1,1,1,1,1}, 0, 1
  fmt.Printf("Input:  %v %d\n", i, k)
  fmt.Printf("Output: %d\n", findPairs(i, k))
  fmt.Printf("Expect: %d\n", o)

  i, k, o = []int{-1,0,1,-2,0}, 2, 2
  fmt.Printf("Input:  %v %d\n", i, k)
  fmt.Printf("Output: %d\n", findPairs(i, k))
  fmt.Printf("Expect: %d\n", o)

  i, k, o = []int{-1,0,1,-2,0}, -1, 0
  fmt.Printf("Input:  %v %d\n", i, k)
  fmt.Printf("Output: %d\n", findPairs(i, k))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(n)
// -- start --

func findPairs(nums []int, k int) int {
  if k < 0 {
    return 0
  }

  m := make(map[int]int)
  for _, num := range nums {
    m[num]++
  }

  c := 0

  for _, num := range nums {
    if k == 0 {
      if m[num] >= 2{
        c++
        m[num] = 0
      }
    } else if m[num+k] > 0 {
      c++
      m[num+k] = 0
    }
  }

  return c
}

// -- end --

