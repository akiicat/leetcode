package main
import "fmt"

func main() {
  i, o := []int{1,2,2,1,1,3}, true
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", uniqueOccurrences(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = []int{1,2}, false
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", uniqueOccurrences(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = []int{-3,0,1,-3,1,1,1,-3,10,0}, true
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", uniqueOccurrences(i))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n)
// M: O(n)
// -- start --

func uniqueOccurrences(arr []int) bool {
  m := make(map[int]int)

  for _, v := range arr {
    m[v]++
  }

  n := make(map[int]bool)
  for _, v := range m {
    _, ok := n[v]
    if ok {
      return false
    }
    n[v] = true
  }

  return true
}

// -- end --

