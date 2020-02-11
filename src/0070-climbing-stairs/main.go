package main
import "math"

// leetcode 509. fibonacci
// https://leetcode.com/articles/climbing-stairs/
// T: O(log(n))
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func climbStairs(n int) int {
  x1, x2 := 1, 1

  for i := 1; i < n; i++ {
    x1, x2 = x2, x1 + x2
  }

  return x2
}

// T: O(log(n))
// M: O(1)
func climbStairsFormula(n int) int {
  sqrt5 := math.Sqrt(5)
  fibn := math.Pow((1+sqrt5)/2, float64(n+1)) - math.Pow((1-sqrt5)/2, float64(n+1))
  return int(math.Round(fibn / sqrt5));
}


// T: O(2 ** n)
// M: O(1)
func climbStairsRecursive(n int) int {
  if n == 0 || n == 1 {
    return 1
  }

  return climbStairsRecursive(n-1) + climbStairsRecursive(n-2)
}


// -- end --

