package main
import "fmt"

func main() {
  i1, i2, o := "abc", "abc", -1
  fmt.Printf("Input:  %s %s\n", i1, i2)
  fmt.Printf("Output: %d\n", findLUSlength(i1, i2))
  fmt.Printf("Expect: %d\n", o)

  i1, i2, o = "abc", "cbc", 3
  fmt.Printf("Input:  %s %s\n", i1, i2)
  fmt.Printf("Output: %d\n", findLUSlength(i1, i2))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(min(a, b))
// M: O(1)
// -- start --

func findLUSlength(a string, b string) int {
  if a == b {
    return -1
  }
  return Max(len(a), len(b))
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

