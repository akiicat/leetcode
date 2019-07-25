package main
import "fmt"

func main() {
  i, o := []string{"hello", "ll"}, 2
  fmt.Printf("Input:  %s, %s\n", i[0], i[1])
  fmt.Printf("Output: %d\n", strStr(i[0], i[1]))
  fmt.Printf("Expect: %d\n", o)

  i, o = []string{"aaaaa", "bba"}, -1
  fmt.Printf("Input:  %s, %s\n", i[0], i[1])
  fmt.Printf("Output: %d\n", strStr(i[0], i[1]))
  fmt.Printf("Expect: %d\n", o)

  i, o = []string{"", ""}, 0
  fmt.Printf("Input:  %s, %s\n", i[0], i[1])
  fmt.Printf("Output: %d\n", strStr(i[0], i[1]))
  fmt.Printf("Expect: %d\n", o)

  i, o = []string{"", "a"}, -1
  fmt.Printf("Input:  %s, %s\n", i[0], i[1])
  fmt.Printf("Output: %d\n", strStr(i[0], i[1]))
  fmt.Printf("Expect: %d\n", o)

  i, o = []string{"a", "a"}, 0
  fmt.Printf("Input:  %s, %s\n", i[0], i[1])
  fmt.Printf("Output: %d\n", strStr(i[0], i[1]))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(N * M)
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

