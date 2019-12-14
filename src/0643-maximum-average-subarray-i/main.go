package main
import "fmt"

func main() {
  i, k, o := []int{1,12,-5,-6,50,3}, 4, 12.75
  fmt.Printf("Input:  %v %d\n", i, k)
  fmt.Printf("Output: %f\n", findMaxAverage(i, k))
  fmt.Printf("Expect: %f\n", o)

  i, k, o = []int{5}, 1, 5.00
  fmt.Printf("Input:  %v %d\n", i, k)
  fmt.Printf("Output: %f\n", findMaxAverage(i, k))
  fmt.Printf("Expect: %f\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func findMaxAverage(nums []int, k int) float64 {
  if len(nums) < k {
    return 0
  }

  sum := 0
  i := 0
  for i < k {
    sum += nums[i]
    i++
  }

  max := sum
  for i < len(nums) {
    sum += nums[i] - nums[i-k]
    if sum > max {
      max = sum
    }
    i++
  }

  return float64(max) / float64(k)
}

// -- end --

