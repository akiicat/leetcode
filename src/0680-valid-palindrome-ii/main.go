package main

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

