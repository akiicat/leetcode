package main
import "math"

// leetcode 70. fibonacci
// https://leetcode.com/articles/climbing-stairs/
// T: O(log(n))
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func fib(n int) int {
  x1, x2 := 0, 1

  for i := 1; i < n; i++ {
    x1, x2 = x2, x1 + x2
  }

  return x2
}

// T: O(log(n))
// M: O(1)
func fibFormula(n int) int {
  sqrt5 := math.Sqrt(5)
  fibn := math.Pow((1+sqrt5)/2, float64(n)) - math.Pow((1-sqrt5)/2, float64(n))
  return int(math.Round(fibn / sqrt5));
}


// T: O(2 ** n)
// M: O(1)
func fibRecursive(n int) int {
  if n == 0 {
    return 0
  }

  if n == 1 {
    return 1
  }

  return fibRecursive(n-1) + fibRecursive(n-2)
}


// -- end --

