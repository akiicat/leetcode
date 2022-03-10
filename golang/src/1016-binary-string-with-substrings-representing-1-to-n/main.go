package main
import "strings"

// T: O(n^2)
// M: O(1)
// -- start --

func queryString(S string, N int) bool {
  for i := N / 2 + 1; i <= N; i++ {
    if strStr(S, itob(i)) < 0 {
      return false
    }
  }
  return true
}

// strconv.FormatInt(int64(i), 2)
func itob(n int) string {
  if n == 0 {
    return "0"
  }

  s := ""
  for n != 0 {
    if n % 2 == 0 {
      s = "0" + s
    } else {
      s = "1" + s
    }
    n /= 2
  }
  return s
}

// leetcode 28. implement strstr
// T: O(n)
// M: O(1)
func strStr(haystack string, needle string) int {
  return strings.Index(haystack, needle)
}

// -- end --

