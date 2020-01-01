package main
import "fmt"

func main() {
  i, o := []int{3,6,1,0}, 1
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", dominantIndex(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,2,3,4}, -1
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", dominantIndex(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func dominantIndex(nums []int) int {
  if len(nums) == 1 {
    return 0
  }

  idx, m1, m2 := 0, 0, 0

  for i, v := range nums {
    if v > m1 {
      idx, m1, m2 = i, v, m1
    } else if v > m2 {
      m2 = v
    }
  }

  if m1 >= m2 << 1 {
    return idx
  }

  return -1
}

// -- end --

