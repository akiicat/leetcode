package main
import "fmt"

func main() {
  i, o := "aba", true
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", validPalindrome(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "abca", true
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", validPalindrome(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "abbca", true
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", validPalindrome(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "abcbca", true
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", validPalindrome(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "abc", false
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", validPalindrome(i))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n)
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

