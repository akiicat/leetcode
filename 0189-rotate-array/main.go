package main
import "fmt"

func main() {
  i, k, o := []int{1,2,3,4,5,6,7}, 3, []int{5,6,7,1,2,3,4}
  fmt.Printf("Input:  %d\n", i)
  rotate(i, k)
  fmt.Printf("Output: %d\n", i)
  fmt.Printf("Expect: %d\n", o)

  i, k, o = []int{-1,-100,3,99}, 2, []int{3,99,-1,-100}
  fmt.Printf("Input:  %d\n", i)
  rotate(i, k)
  fmt.Printf("Output: %d\n", i)
  fmt.Printf("Expect: %d\n", o)

  i, k, o = []int{-1}, 2, []int{-1}
  fmt.Printf("Input:  %d\n", i)
  rotate(i, k)
  fmt.Printf("Output: %d\n", i)
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func rotate(nums []int, k int)  {
  l := len(nums)
  k %= l
  copy(nums, append(nums[l-k:], nums[:l-k]...))
}

// -- end --

