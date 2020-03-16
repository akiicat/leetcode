package main

// T: O(n)
// M: O(n)
// -- start --

func removeDuplicates(S string) string {
  s := []rune{}
  for _, v := range S {
    if len(s) > 0 && s[len(s)-1] == v {
      s = s[:len(s)-1]
    } else {
      s = append(s, v)
    }
  }
  return string(s)
}

// T: O(n)
// M: O(n)
func removeDuplicatesString(S string) string {
  i := 0
  for i < len(S)-1 {
    if S[i] != S[i+1] {
      i++
      continue
    }
    S = S[:i] + S[i+2:]
    i = Max(0, i-1)
  }
  return S
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --
