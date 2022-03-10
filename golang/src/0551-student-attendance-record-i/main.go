package main
import "strings"

// T: O(n)
// M: O(1)
// -- start --

func checkRecord(s string) bool {
  a := 1
  l := 2

  for _, v := range s {
    if v == 'A' {
      if a == 0 {
        return false
      }
      a--
    }

    if v == 'L' {
      if l == 0 {
        return false
      }
      l--
    } else {
      l = 2
    }
  }

  return true
}

func checkRecordStrings(s string) bool {
	return !strings.Contains(s,"LLL") && strings.Count(s,"A") <= 1
}

// -- end --
