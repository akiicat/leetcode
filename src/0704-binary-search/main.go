package main
import "fmt"

func main() {
  i, t, o := []int{-1,0,3,5,9,12}, 9, 4
  fmt.Printf("Input:  %v target=%d\n", i, t)
  fmt.Printf("Output: %d\n", search(i, t))
  fmt.Printf("Expect: %d\n", o)

  i, t, o = []int{-1,0,3,5,9,12}, 2, -1
  fmt.Printf("Input:  %v target=%d\n", i, t)
  fmt.Printf("Output: %d\n", search(i, t))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(log(n))
// M: O(1)
// -- start --

// T: O(log(n))
// M: O(1)
func search(nums []int, target int) int {
  l, m, r := 0, 0, len(nums) - 1

  for l < r {
    m = (l + r) / 2
    if nums[m] == target {
      return m
    }
    if nums[m] > target {
      r = m - 1
    } else {
      l = m + 1
    }
  }

  return -1
}

// T: O(log(n))
// M: O(log(n))
func searchRecursive(nums []int, target int) int {
  return R(nums, target, 0, len(nums) - 1)
}

func R(nums []int, target, l, r int) int {
  if l >= r {
    return -1
  }

  m := (l + r) / 2

  if nums[m] > target {
    return R(nums, target, l, m - 1)
  }

  if nums[m] < target {
    return R(nums, target, m + 1, r)
  }

  return m
}

// -- end --

