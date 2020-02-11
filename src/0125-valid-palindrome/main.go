package main

// T: O(n)
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

