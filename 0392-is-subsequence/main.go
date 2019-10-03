package main
import "fmt"

func main() {
  s, t, o := "abc", "ahbgdc", true
  fmt.Printf("Input:  %s %s\n", s, t)
  fmt.Printf("Output: %t\n", isSubsequence(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "axc", "ahbgdc", false
  fmt.Printf("Input:  %s %s\n", s, t)
  fmt.Printf("Output: %t\n", isSubsequence(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "aec", "abcde", false
  fmt.Printf("Input:  %s %s\n", s, t)
  fmt.Printf("Output: %t\n", isSubsequence(s, t))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func isSubsequence(s string, t string) bool {
  i, j := 0, 0

  for i < len(s) && j < len(t) {
    if s[i] == t[j] {
      i++
    }
    j++
  }

  return i == len(s)
}

func isSubsequenceEarlyStop(s string, t string) bool {
  p := 0

  ns, nt := len(s), len(t)

  if ns == 0 {
    return true
  }

  if nt == 0 {
    return false
  }

  for _, v := range t {
    if s[p] == byte(v) {
      p++

      if p >= ns {
        return true
      }
    }
  }

  return false
}

// -- end --
