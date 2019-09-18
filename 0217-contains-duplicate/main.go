package main
import "fmt"

func main() {
  i, o := []int{1,2,3,1}, true
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", containsDuplicate(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = []int{1,2,3,4}, false
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", containsDuplicate(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = []int{1,1,1,3,3,4,3,2,4,2}, true
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", containsDuplicate(i))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n)
// M: O(n)
// -- start --

func containsDuplicate(nums []int) bool {
  m := make(map[int]bool)

  for _, v := range nums {
    _, ok := m[v]
    if ok {
      return true
    } else {
      m[v] = true
    }
  }

  return false
}

// -- end --

