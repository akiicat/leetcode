package main
import "fmt"

func main() {
  i, o := "abccccdd", 7
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", longestPalindrome(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func longestPalindrome(s string) int {
  m := [58]int{}

  for _, v := range s {
    m[v-'A']++
  }

  sum, odd := 0, 0

  for _, v := range m {
    sum += v
    if v & 0x1 == 1 {
      sum--
      odd = 1
    }
  }

  return sum + odd
}

// -- end --

