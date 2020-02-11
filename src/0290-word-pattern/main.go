package main
import "strings"

// leetcode 205.
// T: O(n)
// M: O(c) c for len(pattern)
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

