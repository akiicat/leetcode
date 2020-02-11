package main
import "strings"

// T: O(n + m) n, m are the lengths of A, B respectively
// M: O(n + m)
// -- start --

func uncommonFromSentences(A string, B string) []string {
  m := make(map[string]int)

  for _, v := range strings.Split(A, " ") {
    m[v]++
  }
  for _, v := range strings.Split(B, " ") {
    m[v]++
  }

  rtn := []string{}
  for k, v := range m {
    if v == 1 {
      rtn = append(rtn, k)
    }
  }

  return rtn
}

// -- end --

