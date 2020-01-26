package main
import "fmt"

func main() {
  i, o := "RLRRLLRLRL", 4
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", balancedStringSplit(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "RLLLLRRRLR", 3
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", balancedStringSplit(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "LLLLRRRR", 1
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", balancedStringSplit(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "RLRRRLLRLL", 2
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", balancedStringSplit(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func balancedStringSplit(s string) int {

  count, b := 0, 0

  for _, v := range s {
    switch v {
    case 'R': b++
    case 'L': b--
    }
    if b == 0 {
      count++
    }
  }

  return count
}

// -- end --

