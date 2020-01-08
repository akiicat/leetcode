package main
import "fmt"
import "strings"

func main() {
  a, b, o := "abcde", "cdeab", true
  fmt.Printf("Input:  %s %s\n", a, b)
  fmt.Printf("Output: %t\n", rotateString(a, b))
  fmt.Printf("Expect: %t\n", o)

  a, b, o = "abcde", "abced", false
  fmt.Printf("Input:  %s %s\n", a, b)
  fmt.Printf("Output: %t\n", rotateString(a, b))
  fmt.Printf("Expect: %t\n", o)
}

// https://leetcode.com/articles/rotate-string/
// leetcode 28. 459.
// T: O(n + k)
// M: O(k)
// -- start --

func rotateString(A string, B string) bool {
  if len(A) != len(B) {
    return false
  }

  if strStr(A + A, B) == -1 {
    return false
  }

  return true
}

// leetcode 28.
// KMP Shift Table
// T: O(n + k)
// M: O(k)
func strStr(haystack string, needle string) int {
  if len(needle) == 0 {
    return 0
  }

  // haystack[i]
  // needle[j]

  prefix := make([]int, len(needle) + 1) // prefix table
  for i := range prefix {
    prefix[i] = 1
  }

  // create shift table
  l := -1 // the len of the maximum same prefix
  for j := 0; j < len(needle); j++ {
    for l >= 0 && needle[l] != needle[j] {
      l = l - prefix[l]
    }
    prefix[j+1] = j - l
    l++
  }

  // search by shift table
  j := 0
  for i := range haystack {
    for j >= 0 && haystack[i] != needle[j] {
      j = j - prefix[j]
    }

    if j == len(needle)-1 {
      return i - j
    }

    j++
  }

  return -1
}

// T: O(n ** 2)
// M: O(n)
func rotateStringSimpleCheck(A string, B string) bool {
  return len(A) == len(B) && strings.Contains(A + A, B)
}

// -- end --

