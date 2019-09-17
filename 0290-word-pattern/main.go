package main
import "fmt"
import "strings"

func main() {
  s, t, o := "abba", "dog cat cat dog", true
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", wordPattern(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "abba", "dog cat cat fish", false
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", wordPattern(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "aaaa", "dog cat cat dog", false
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", wordPattern(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "abba", "dog dog dog dog", false
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", wordPattern(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "aaa", "dog dog dog dog", false
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", wordPattern(s, t))
  fmt.Printf("Expect: %t\n", o)
}

// leetcode 205.
// T: O(n)
// M: O(n)
// -- start --

func wordPattern(pattern string, str string) bool {
  m1 := make(map[byte]string)
  m2 := make(map[string]byte)

  strs := strings.Split(str, " ")

  if len(pattern) != len(strs) {
    return false
  }

  for i := 0; i < len(pattern); i++ {
    v1, ok1 := m1[pattern[i]]
    v2, ok2 := m2[strs[i]]

    if !ok1 {
      m1[pattern[i]] = strs[i]
    }
    if !ok2 {
      m2[strs[i]] = pattern[i]
    }

    if ok1 && v1 != strs[i] || ok2 && v2 != pattern[i] {
      return false
    }
  }

  return true
}

// -- end --

