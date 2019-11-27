package main
import "fmt"

func main() {
  i, o := []int{1,3,2,2,5,2,3,7}, 5
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", findLHS(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,1,1,1}, 0
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", findLHS(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(n)
// -- start --

func findLHS(nums []int) int {
  m := make(map[int]int)

  for _, v := range nums {
    m[v]++
  }

  max := 0
  for k, v := range m {
    next, ok := m[k+1]
    if ok && v + next > max {
      max = v + next
    }
  }

  return max
}

// -- end --

