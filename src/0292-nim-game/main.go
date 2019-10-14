package main
import "fmt"

func main() {
  i, o := 0, false
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", canWinNim(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 1, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", canWinNim(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 2, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", canWinNim(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 3, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", canWinNim(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 4, false
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", canWinNim(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 5, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", canWinNim(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 6, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", canWinNim(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 7, true
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", canWinNim(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = 8, false
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %t\n", canWinNim(i))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(1)
// M: O(1)
// -- start --

func canWinNim(n int) bool {
  return n % 4 != 0
}

// -- end --

