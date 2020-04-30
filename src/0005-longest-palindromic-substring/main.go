package main
import "log"

// T: O(n)
// M: O(n)
// -- start --

// Manacher's Algorithm
// https://www.hackerrank.com/topics/manachers-algorithm
// T: O(n)
// M: O(n)
func longestPalindrome(s string) string {
  t := preProcess(s)
  n := len(t)
  p := make([]int, n)

  c, r := 0, 0
  for i := 1; i < n-1; i++ {
    mirror := 2 * c - i

    if r > i {
      p[i] = Min(r - i, p[mirror])
    } else {
      p[i] = 0
    }

    for t[i + 1 + p[i]] == t[i - 1 - p[i]] {
      p[i]++
    }

    if i + p[i] > r {
      c = i
      r = i + p[i]
    }
  }

  max := 0
  for i := 1; i < n - 1; i++ {
    if p[i] > p[max] {
      max = i
    }
  }

  return s[(max - p[max]) / 2:(max + p[max]) / 2]
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

// T: O(n ** 2)
// M: O(1)
func longestPalindromeExpandAroundCenter(s string) string {
  if len(s) == 0 {
    return ""
  }

  l, r := 0, 0
  for i := 0; i < len(s); i++ {
    n1 := expandAroundCenter(s, i, i)
    n2 := expandAroundCenter(s, i, i+1)

    n := Max(n1, n2)

    if n > r - l {
      l, r = i - (n - 1) / 2, i + n / 2
    }
  }

  return s[l:r+1]
}

func expandAroundCenter(s string, l, r int) int {
  for l >= 0 && r < len(s) && s[l] == s[r] {
    l--
    r++
  }

  return r - l - 1
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// T: O(n ** 3)
// M: O(1)
func longestPalindromeBruteForce(s string) string {
  for l := len(s); l > 0; l-- {
    for i := 0; i < len(s)-l+1; i++ {
      if isPalindrome(s[i:i+l]) {
        return s[i:i+l]
      }
    }
  }
  return ""
}

func isPalindrome(s string) bool {
  left, right := 0, len(s) - 1

  for left < right {
    if s[left] != s[right] {
      return false
    }

    left++
    right--
  }

  return true
}

// -- end --

