package main
import "fmt"

func main() {
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: true\n", "aba", validPalindrome("aba"))
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: true\n", "abca", validPalindrome("abca"))
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: true\n", "abbca", validPalindrome("abbca"))
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: true\n", "abcbca", validPalindrome("abcbca"))
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: false\n", "abc", validPalindrome("abc"))
}

// T: O(N)
// M: O(1)
// -- start --

func validPalindrome(s string) bool {
  left, right := 0, len(s) - 1

  for left < right {
    if s[left] != s[right] {
      return judge(s, left+1, right) || judge(s, left, right-1)
    }

    left++
    right--
  }

  return true
}

func judge(s string, left, right int) bool {
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

