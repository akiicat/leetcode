package main

// T: O(n)
// M: O(1)
// -- start --

// Linear Scan
// T: O(n)
// M: O(1)
func countBinarySubstrings(s string) int {
  pre, cur := s[0], 1
  sum, res := 0, 0
  for i := 1; i < len(s); i++ {
    if pre == s[i] {
      cur++
      continue
    }

    sum += Min(cur, res)
    pre, cur, res = s[i], 1, cur
  }
  sum += Min(cur, res)

  return sum
}


// Group By Character
// T: O(n)
// M: O(n)
func countBinarySubstringsArray(s string) int {
  m := []int{}

  p, c := s[0], 1
  for i := 1; i < len(s); i++ {
    if p != s[i] {
      m = append(m, c)
      p, c = s[i], 0
    }

    c++
  }
  m = append(m, c)


  sum := 0
  for i := 1; i < len(m); i++ {
    sum += Min(m[i-1], m[i])
  }

  return sum
}

func Min(a, b int) int {
  if a > b {
    return b
  }
  return a
}

// -- end --

