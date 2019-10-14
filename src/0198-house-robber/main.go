package main
import "fmt"

func main() {
  i, o := []int{1,2,3,1}, 4
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", rob(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{2,7,9,3,1}, 12
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", rob(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{114,117,207,117,235,82,90,67,143,146,53,108,200,91,80,223,58,170,110,236,81,90,222,160,165,195,187,199,114,235,197,187,69,129,64,214,228,78,188,67,205,94,205,169,241,202,144,240}, 4173
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", rob(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --


// T: O(n)
// M: O(1)
func rob(nums []int) int {
  last, result := 0, 0
  for _, v := range nums {
    tmp := result
    result = Max(last + v, result)
    last = tmp
  }
  return result
}


// T: O(n ** 2)
// M: O(n ** 2)
func robRecursive(nums []int) int {
  if len(nums) == 0 {
    return 0
  }

  if len(nums) == 1 {
    return nums[0]
  }

  // take: nums[0] + rob(nums[2:])
  // not take: rob(nums[1:])
  return Max(nums[0] + rob(nums[2:]), rob(nums[1:]))
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

