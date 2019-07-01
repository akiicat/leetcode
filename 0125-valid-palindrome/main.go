package main
import "fmt"

func main() {
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: true\n", "A man, a plan, a canal: Panama", isPalindrome("A man, a plan, a canal: Panama"))
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: false\n", "race a car", isPalindrome("race a car"))
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: false\n", "OP", isPalindrome("OP"))
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: false\n", "0P", isPalindrome("0P"))
}

// T: O(N)
// M: O(1)
// -- start --

func isPalindrome(s string) bool {
  left, right := 0, len(s) - 1

  for left < right {

    for left < right && isValidChr(s[left]) {
      left++
    }

    for left < right && isValidChr(s[right]) {
      right--
    }

    if getByte(s[left]) != getByte(s[right]) {
      return false
    }

    left++
    right--
  }

  return true
}

func isValidChr(r byte) bool {
  return !(48 <= r && r <= 57 || 97 <= r && r <= 122 || 65 <= r && r <= 90)
}

func getByte(r byte) byte {
  if 65 <= r && r <= 90 {
    return r + 32
  }
  return r
}

// -- end --

