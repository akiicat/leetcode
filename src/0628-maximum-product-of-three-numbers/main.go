package main
import "fmt"

func main() {
  i, o := []int{1,2,3}, 6
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", maximumProduct(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,2,3,4}, 24
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", maximumProduct(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{-1,2,-3,4}, 12
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", maximumProduct(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{-1,-2,-3}, -6
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", maximumProduct(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{-4,-3,-2,-1,60}, 720
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", maximumProduct(i))
  fmt.Printf("Expect: %d\n", o)
}

// leetcode 414. 506.
// T: O(n)
// M: O(1)
// -- start --

func maximumProduct(nums []int) int {
  if len(nums) == 3 {
    return nums[0] * nums[1] * nums[2]
  }

  max := [3]int{}
  min := [2]int{}

  for _, v := range nums {
    if v > 0 {
      if v > max[0] {
        max[0], max[1], max[2] = v, max[0], max[1]
      } else if v > max[1] {
        max[1], max[2] = v, max[1]
      } else if v > max[2] {
        max[2] = v
      }
    } else {
      if v < min[0] {
        min[0], min[1] = v, min[0]
      } else if v < min[1] {
        min[1] = v
      }
    }
  }

  return Max(max[0] * max[1] * max[2], max[0] * min[0] * min[1])
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

