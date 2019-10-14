package main
import "fmt"

func main() {
  i, o := []int{10,15,20}, 15
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", minCostClimbingStairs(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,100,1,1,1,100,1,1,100,1}, 6
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", minCostClimbingStairs(i))
  fmt.Printf("Expect: %d\n", o)
}

// https://leetcode.com/articles/min-cost-climbing-stairs/
// T: O(n)
// M: O(1)
// -- start --

func minCostClimbingStairs(cost []int) int {
  return minCostClimbingStairsDynamicProgramming(cost)
  // return minCostClimbingStairsRecursive(append(cost, 0), len(cost))
}

// T: O(n)
// M: O(1)
func minCostClimbingStairsDynamicProgramming(cost []int) int {
  r1, r2 := cost[0], cost[1]
  for i := 2; i < len(cost); i++ {
    r := cost[i] + min(r1, r2)
    r1, r2 = r2, r
  }

  return min(r1, r2)
}

// T: O(2 ** n)
// M: O(1)
func minCostClimbingStairsRecursive(cost []int, n int) int {
  if n < 0 {
    return 0
  }

  return cost[n] + min(minCostClimbingStairsRecursive(cost, n-1), minCostClimbingStairsRecursive(cost, n-2))
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

