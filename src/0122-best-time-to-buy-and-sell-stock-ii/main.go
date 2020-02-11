package main

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

