package main
import "fmt"

func main() {
  i, o := 2, 2
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", climbStairs(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 3, 3
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", climbStairs(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 4, 5
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", climbStairs(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 5, 8
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", climbStairs(i))
  fmt.Printf("Expect: %d\n", o)
}

// fibonacci
// T: O(N)
// M: O(1)
// -- start --

// T: O(N)
// M: O(1)
func climbStairs(n int) int {
  x1, x2 := 1, 1

  for i := 1; i < n; i++ {
    x1, x2 = x2, x1 + x2
  }

  return x2
}

// T: O(N ** 2)
// M: O(1)
func climbStairsRecursive(n int) int {
  if n == 0 || n == 1 {
    return 1
  }

  return climbStairs(n-1) + climbStairs(n-2)
}


// -- end --

