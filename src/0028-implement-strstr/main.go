package main
import "fmt"

func main() {
  h, n, o := "hello", "ll", 2
  fmt.Printf("Input:  %s, %s\n", h, n)
  fmt.Printf("Output: %d\n", strStr(h, n))
  fmt.Printf("Expect: %d\n", o)

  h, n, o = "aaaaa", "bba", -1
  fmt.Printf("Input:  %s, %s\n", h, n)
  fmt.Printf("Output: %d\n", strStr(h, n))
  fmt.Printf("Expect: %d\n", o)

  h, n, o = "", "", 0
  fmt.Printf("Input:  %s, %s\n", h, n)
  fmt.Printf("Output: %d\n", strStr(h, n))
  fmt.Printf("Expect: %d\n", o)

  h, n, o = "", "a", -1
  fmt.Printf("Input:  %s, %s\n", h, n)
  fmt.Printf("Output: %d\n", strStr(h, n))
  fmt.Printf("Expect: %d\n", o)

  h, n, o = "a", "a", 0
  fmt.Printf("Input:  %s, %s\n", h, n)
  fmt.Printf("Output: %d\n", strStr(h, n))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func strStr(haystack string, needle string) int {
  hLen, nLen := len(haystack), len(needle)

  if nLen == 0 {
    return 0
  }

  for i := 0; i < hLen - nLen + 1; i++ {
    if haystack[i:i+nLen] == needle {
      return i
    }
  }

  return -1
}

// -- end --

