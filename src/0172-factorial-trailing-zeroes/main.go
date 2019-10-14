package main
import "fmt"

func main() {
  i, o := 3, 0
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", trailingZeroes(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 5, 1
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", trailingZeroes(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 25, 6
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", trailingZeroes(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(log(n))
// M: O(1)
// -- start --


func trailingZeroes(n int) int {
  sum := 0

  for {
    n /= 5
    sum += n
    if n == 0 {
      break
    }
  }

  return sum
}

// -- end --

