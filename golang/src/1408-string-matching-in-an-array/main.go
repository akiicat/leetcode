package main
import "sort"
import "strings"

// T: O(n ** 2)
// M: O(1)
// -- start --

func stringMatching(words []string) []string {
  sort.Slice(words, func(i, j int) bool {
    return len(words[i]) < len(words[j])
  })

  rtn := []string{}
  for len(words) != 0 {
    s := words[0]
    words = words[1:]
    for _, v := range words {
      if (strStr(v, s) != -1) {
        rtn = append(rtn, s)
        break
      }
    }
  }

  return rtn
}

// leetcode 28.
func strStr(haystack string, needle string) int {
  return strings.Index(haystack, needle)
}

// -- end --

