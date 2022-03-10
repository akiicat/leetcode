package main
import "log"

// T: O(n)
// M: O(n)
// -- start --

// leetcode 5.
// Manacher's Algorithm
// https://www.hackerrank.com/topics/manachers-algorithm
// T: O(n)
// M: O(n)
func countSubstrings(s string) int {
  t := preProcess(s)
  p := make([]int, len(t))

  r, c := 0, 1
  for i := 1; i < len(t)-1; i++ {
    mirror := 2 * c - i

    if i < r {
      p[i] = Min(r - i, p[mirror])
    } else {
      p[i] = 0
    }

    for t[i-p[i]-1] == t[i+p[i]+1] {
      p[i]++
    }

    if r < i + p[i] {
      c = i
      r = i + p[i]
    }
  }

  sum := 0
  for _, v := range p {
    sum += (v+1)/2
  }

  return sum
}

func preProcess(s string) string {
  if len(s) == 0 {
    return "^$"
  }
  rtn := "^#"
  for _, v := range s {
    rtn += string([]rune{v, '#'})
  }
  return rtn + "$"
}

func Min(a, b int) int {
  if a > b {
    return b
  }
  return a
}


// T: O(n ^ 2)
// M: O(1)
func countSubstringsExpandAroundCenter(s string) int {
  sum := 0
  for i := range s {
    sum += expand(s, i, i)
    sum += expand(s, i, i+1)
  }
  return sum
}

func expand(s string, l, r int) int {
  sum := 0
  for l >= 0 && r < len(s) && s[l] == s[r] {
    l--
    r++
    sum++
  }
  return sum
}

// -- end --

