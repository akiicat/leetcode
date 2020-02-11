package main
import "strings"

// T: O(n + m)
// M: O(1)
// -- start --

// T: O(n + m)
// M: O(1)
func repeatedStringMatch(A string, B string) int {
  bound, cur, str := len(B) / len(A) + 2, 1, A

  for cur <= bound {
    if strings.Index(A, B) != -1 {
      return cur
    }

    cur, A = cur + 1, A + str
  }

  return -1
}

// -- end --

