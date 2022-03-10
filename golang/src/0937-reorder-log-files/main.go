package main
import "sort"

// T: O(n * log(n))
// M: O(n)
// -- start --

func reorderLogFiles(logs []string) []string {
  sort.SliceStable(logs, func(i, j int) bool {
    ia, a := getStr(logs[i])
    ib, b := getStr(logs[j])

    if isDigit(a) && isDigit(b) {
      return i < j
    }

    if isDigit(b) {
      return true
    }

    if isDigit(a) {
      return false
    }

    if a == b {
      return ia < ib
    }

    return a < b
  })

  return logs
}

func isDigit(s string) bool {
  return '0' <= s[0] && s[0] <= '9'
}

func getStr(s string) (string, string) {
  for i, v := range s {
    if v == ' ' {
      return s[:i], s[i+1:]
    }
  }
  return "", ""
}

// -- end --

