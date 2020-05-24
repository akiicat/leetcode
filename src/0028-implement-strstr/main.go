package main
import "strings"

// leetcode 459. 796.
// T: O(n + k)
// M: O(k)
// -- start --

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

// KMP Prefix Table
// T: O(n + k)
// M: O(k)
func strStrKmpWithPrefixTable(haystack string, needle string) int {
  if len(needle) == 0 {
    return 0
  }

  // haystack[i]
  // needle[j]

  prefix := make([]int, len(needle) + 1) // prefix table
  prefix[0] = -1

  // create prefix table
  l := 0 // the len of the maximum same prefix
  for j := 1; j < len(needle); j++ {
    for l > 0 && needle[l] != needle[j] {
      l = prefix[l]
    }

    if needle[l] == needle[j] {
      l++
    }

    prefix[j+1] = l
  }

  // search by prefix table
  j := 0
  for i := range haystack {
    for j != -1 && haystack[i] != needle[j] {
      j = prefix[j]
    }

    if j == len(needle)-1 {
      return i - j
    }

    i++
    j++
  }

  return -1
}

// T: O(n * k)
// M: O(1)
func strStrBruteForce(haystack string, needle string) int {
  hLen, nLen := len(haystack), len(needle)

  if nLen == 0 {
    return 0
  }

  for i := 0; i < hLen - nLen + 1; i++ {
    if haystack[i:i+nLen] == needle {
      return i
    }
  }

  return -1
}

// T: O(n + k)
// M: O(1)
func strStrLibrary(haystack string, needle string) int {
  return strings.Index(haystack, needle)
}


// -- end --

