package main
import "fmt"

func main() {
  i, o := []int{7,1,5,3,6,4}, 5
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

// leetcode 53.
// https://leetcode.com/articles/best-time-to-buy-and-sell-stock/
// T: O(n)
// M: O(1)
// -- start --

func maxProfit(prices []int) int {
  return maxProfitOnePass(prices []int)
}

// T: O(n)
// M: O(1)
func maxProfitOnePass(prices []int) int {
  m := 0 // max
  low := 1<<31

  for _, v := range prices {
    low = min(low, v)
    m = max(v - low, m)
  }

  return m
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// T: O(n)
// M: O(1)
func maxProfitSubArray(prices []int) int {
  if len(prices) == 0 {
    return 0
  }

  for i := 0; i < len(prices) - 1; i++ {
    prices[i] = prices[i+1] - prices[i]
  }

  prices[len(prices) - 1] = 0

  return max(maxSubArray(prices), 0)
}

// leetcode 53.
func maxSubArray(nums []int) int {
  m, cur := nums[0], 0

  for _, v := range nums {
    cur = max(cur+v, v)
    m = max(cur, m)
  }

  return m // max
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

