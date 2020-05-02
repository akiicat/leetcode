package main

// T: O(n^2)
// M: O(n)
// -- start --

// Dynamic Programming
// https://www.geeksforgeeks.org/longest-palindrome-subsequence-space/
// https://www.youtube.com/watch?v=_nCsPn7_OgI
// https://www.youtube.com/watch?v=i1o6rlqiSZM
// T: O(n^2)
// M: O(n)
func longestPalindromeSubseq(s string) int {
  m0 := make([]int, len(s)) // 0 if length == 0; even default value
  m1 := make([]int, len(s)) // 1 if length == 1; odd default value
  for i := range m1 { // i == j
    m1[i] = 1
  }
  m2 := make([]int, len(s)) // temp
  for l := 2; l <= len(s); l++ { // length from 2 to len(s)
    for i := 0; i <= len(s)-l; i++ { // from left to right
      if s[i] == s[i+l-1] {
        m2[i] = m0[i+1] + 2
      } else {
        m2[i] = Max(m1[i], m1[i+1])
      }
    }
    m0, m1, m2 = m1, m2, m0
  }
  return m1[0]
}


// T: O(n^2)
// M: O(n^2)
func longestPalindromeSubseqCacheRecursive(s string) int {
  m := make([][]int, len(s))
  for i := range m {
    m[i] = make([]int, len(s))
  }
  return R2(s, m, 0, len(s)-1)
}

func R2(s string, m [][]int, l, r int) int {
  if l > r {
    return 0
  }
  if l == r {
    return 1
  }
  // do not repeatly calc.
  if m[l][r] != 0 {
    return m[l][r]
  }

  if s[l] == s[r] {
    m[l][r] = R2(s, m, l+1, r-1) + 2
  } else {
    m[l][r] = Max(R2(s, m, l+1, r), R2(s, m, l, r-1))
  }

  return m[l][r]
}

// T: O(n^n)
// M: O(n)
func longestPalindromeSubseqRecursive(s string) int {
  return R1(s, 0, len(s)-1)
}

func R1(s string, l, r int) int {
  if l > r {
    return 0
  }
  if l == r {
    return 1
  }

  m := 0
  if s[l] == s[r] {
    m = R1(s, l+1, r-1) + 2
  } else {
    m1 := R1(s, l+1, r)
    m2 := R1(s, l, r-1)
    m = Max(m1, m2)
  }

  return m
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

