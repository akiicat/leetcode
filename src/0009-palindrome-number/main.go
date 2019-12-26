package main
import "fmt"

func main() {
  i, o := 0, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", isPalindrome(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 1, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", isPalindrome(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 121, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", isPalindrome(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = -121, false
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", isPalindrome(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 313, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", isPalindrome(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 1122, false
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", isPalindrome(i))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func isPalindrome(x int) bool {
  if x < 0 {
    return false
  }

  y := 0
  tmp := x
  for tmp != 0 {
    y = 10 * y + tmp % 10
    tmp /= 10
  }

  return x == y
}

// -- end --

