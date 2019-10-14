package main
import "fmt"

func main() {
  i, o = 4, 4
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", tribonacci(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 4, 4
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", tribonacci(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 25, 1389537
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", tribonacci(i))
  fmt.Printf("Expect: %d\n", o)
}

// Formula
// http://mathworld.wolfram.com/TribonacciNumber.html
// T: O(log(n))
// M: O(1)
// -- start --

func tribonacci(n int) int {
  return tribonacciDynamicProgramming(n)
}

// T: O(n)
// M: O(1)
func tribonacciDynamicProgramming(n int) int {
  if n == 0 {
    return 0
  }

  if n == 1 || n == 2 {
    return 1
  }

  r1, r2, r3 := 0, 1, 1

  for i := 0; i < n - 2; i++ {
    r1, r2, r3 = r2, r3, r1 + r2 + r3
  }

  return r3
}

// T: O(3 ** n)
// M: O(1)
func tribonacciRecursive(n int) int {
  if n == 0 {
    return 0
  }

  if n == 1 || n == 2 {
    return 1
  }

  return tribonacciRecursive(n-1) +
         tribonacciRecursive(n-2) +
         tribonacciRecursive(n-3)
}

// -- end --

