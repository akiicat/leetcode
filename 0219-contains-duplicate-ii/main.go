package main
import "fmt"

func main() {
  i, k, o := []int{1,2,3,1}, 3, true
  fmt.Printf("Input:  %v, %d\n", i, k)
  fmt.Printf("Output: %t\n", containsNearbyDuplicate(i, k))
  fmt.Printf("Expect: %t\n", o)

  i, k, o = []int{1,0,1,1}, 1, true
  fmt.Printf("Input:  %v, %d\n", i, k)
  fmt.Printf("Output: %t\n", containsNearbyDuplicate(i, k))
  fmt.Printf("Expect: %t\n", o)

  i, k, o = []int{1,2,3,1,2,3}, 2, false
  fmt.Printf("Input:  %v, %d\n", i, k)
  fmt.Printf("Output: %t\n", containsNearbyDuplicate(i, k))
  fmt.Printf("Expect: %t\n", o)
}

// leetcode 217.
// T: O(n)
// M: O(n)
// -- start --

func containsNearbyDuplicate(nums []int, k int) bool {
  m := make(map[int]int)

  for i, v := range nums {
    j, ok := m[v]
    if ok && i - j <= k {
      return true
    } else {
      m[v] = i
    }
  }

  return false
}

// -- end --

