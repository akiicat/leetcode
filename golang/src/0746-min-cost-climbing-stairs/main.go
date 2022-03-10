package main

// https://leetcode.com/articles/min-cost-climbing-stairs/
// T: O(n)
// M: O(1)
// -- start --

// Dynamic Programming
// T: O(n)
// M: O(1)
func minCostClimbingStairs(cost []int) int {
  r1, r2 := cost[0], cost[1]
  for i := 2; i < len(cost); i++ {
    r := cost[i] + min(r1, r2)
    r1, r2 = r2, r
  }

  return min(r1, r2)
}

func minCostClimbingStairsRecursive(cost []int) int {
  return R(append(cost, 0), len(cost))
}

// T: O(2 ** n)
// M: O(1)
func R(cost []int, n int) int {
  if n < 0 {
    return 0
  }

  return cost[n] + min(R(cost, n-1), R(cost, n-2))
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

