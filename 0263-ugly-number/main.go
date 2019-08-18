package main
import "fmt"

func main() {
  i, o := 0, false
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", isUgly(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 1, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", isUgly(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 6, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", isUgly(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 8, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", isUgly(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 14, false
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", isUgly(i))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(1)
// M: O(1)
// -- start --

// T: O(n)
// M: O(n)
func isUgly(num int) bool {
  if num < 1 {
    return false
  }

  for num % 2 == 0 {
    num /= 2
  }

  for num % 3 == 0 {
    num /= 3
  }

  for num % 5 == 0 {
    num /= 5
  }


  if num == 1 {
    return true
  }

  return false
}

// -- end --
