package main
import "fmt"

func main() {
  s, t, o := "egg", "add", true
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", isIsomorphic(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "foo", "bar", false
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", isIsomorphic(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "paper", "title", true
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", isIsomorphic(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "ab", "aa", false
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", isIsomorphic(s, t))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n)
// M: O(1) not exceed 26 chars
// -- start --

func isIsomorphic(s string, t string) bool {
  m1 := make(map[byte]byte)
  m2 := make(map[byte]byte)

  for i := 0; i < len(s); i++ {
    v1, ok1 := m1[s[i]]
    if !ok1 {
      m1[s[i]] = t[i]
    } else if v1 != t[i] {
      return false
    }

    v2, ok2 := m2[t[i]]
    if !ok2 {
      m2[t[i]] = s[i]
    } else if v2 != s[i] {
      return false
    }
  }

  return true
}

// -- end --

