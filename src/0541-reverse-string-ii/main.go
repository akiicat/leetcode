package main

// T: O(n) n is the length of the string
// M: O(1)
// -- start --

func reverseStr(s string, k int) string {
  n := len(s)
  if n < 2 {
    return s
  }

  b := []byte(s)
  for i := 0; i < n; i += k << 1 {
    Reverse(b, i, Min(i+k-1, n-1))
  }

  return string(b)
}

// leetcode 0344-reverse-string
func Reverse(b []byte, l, r int) {
  for r > l {
    b[l], b[r] = b[r], b[l]
    l, r = l+1, r-1
  }
  return
}

func Min(a, b int) int {
  if a > b {
    return b
  }
  return a
}

// -- end --

