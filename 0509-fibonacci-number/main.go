package main
import "fmt"
import "math"

func main() {
  i, o := 2, 1
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", fib(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 3, 2
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", fib(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 4, 3
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", fib(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 5, 5
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", fib(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 9, 34
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", fib(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 10, 55
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", fib(i))
  fmt.Printf("Expect: %d\n", o)
}

// leetcode 70. fibonacci
// https://leetcode.com/articles/climbing-stairs/
// T: O(log(N))
// M: O(1)
// -- start --

// T: O(N)
// M: O(1)
func fib(n int) int {
  x1, x2 := 0, 1

  for i := 1; i < n; i++ {
    x1, x2 = x2, x1 + x2
  }

  return x2
}

// T: O(log(N))
// M: O(1)
func fibFormula(n int) int {
  sqrt5 := math.Sqrt(5)
  fibn := math.Pow((1+sqrt5)/2, float64(n)) - math.Pow((1-sqrt5)/2, float64(n))
  return int(math.Round(fibn / sqrt5));
}


// T: O(2 ** N)
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

