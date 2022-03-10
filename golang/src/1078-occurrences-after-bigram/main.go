package main
import "strings"

// T: O(n) n is the number of words
// M: O(1)
// -- start --

func findOcurrences(text string, first string, second string) []string {
  t := strings.Split(text, " ")

  if len(t) < 3 {
    return []string{}
  }

  rtn := []string{}

  for i := 0; i < len(t) - 2; i++ {
    if t[i] == first && t[i+1] == second {
      rtn = append(rtn, t[i+2])
    }
  }

  return rtn
}

// -- end --

