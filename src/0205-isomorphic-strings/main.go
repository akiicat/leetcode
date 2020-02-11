package main

// leetcode 290.
// T: O(n)
// M: O(1) not exceed 26 chars
// -- start --

func isIsomorphic(s string, t string) bool {
  m1 := make(map[byte]byte)
  m2 := make(map[byte]byte)

  for i := 0; i < len(s); i++ {
    v1, ok1 := m1[s[i]]
    v2, ok2 := m2[t[i]]

    if !ok1 {
      m1[s[i]] = t[i]
    }
    if !ok2 {
      m2[t[i]] = s[i]
    }

    if ok1 && v1 != t[i] || ok2 && v2 != s[i] {
      return false
    }
  }

  return true
}

// -- end --

