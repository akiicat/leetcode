package main
import "fmt"

func main() {
  i, o := []int{7,1,5,3,6,4}, 7
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", maxProfit(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,2,3,4,5}, 4
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", maxProfit(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{7,6,4,3,1}, 0
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", maxProfit(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{}, 0
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", maxProfit(i))
  fmt.Printf("Expect: %d\n", o)
}

// https://leetcode.com/articles/best-time-to-buy-and-sell-stock-ii/
// T: O(n)
// M: O(1)
// -- start --

// Single One Pass
func maxProfit(prices []int) int {
  sum := 0

  // calcuate the increasing part
  for i := 0; i < len(prices) - 1; i++ {
    diff := prices[i+1] - prices[i]
    if diff > 0 {
      sum += diff
    }
  }

  return sum
}

// -- end --

