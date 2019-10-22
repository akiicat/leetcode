package main
import "fmt"

func main() {
  i, o := []int{1,1,0,1,1,1}, 3
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", findMaxConsecutiveOnes(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,0,1,1,0,1}, 2
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", findMaxConsecutiveOnes(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func findMaxConsecutiveOnes(nums []int) int {
  n := len(nums)
  p, q, max := 0, 0, 0

  for q < n {

    // skip zeros
    p = q
    for p < n && nums[p] == 0 {
      p++
    }

    // skip ones
    q = p
    for q < n && nums[q] == 1 {
      q++
    }

    if q - p > max {
      max = q - p
    }
  }

  return max
}

func findMaxConsecutiveOnesCounter(nums []int) int {
  c, max := 0, 0

  for _, v := range nums {
    if v == 1 {
      c++
      continue
    }

    if c > max {
      max = c
    }

    c = 0
  }

  if c > max {
    return c
  }

  return max
}

// -- end --

