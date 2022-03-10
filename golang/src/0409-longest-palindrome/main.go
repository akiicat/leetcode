package main

// T: O(n)
// M: O(1)
// -- start --

func longestPalindrome(s string) int {
  m := [58]int{}

  for _, v := range s {
    m[v-'A']++
  }

  sum, odd := 0, 0

  for _, v := range m {
    sum += v
    if v & 0x1 == 1 {
      sum--
      odd = 1
    }
  }

  return sum + odd
}

// -- end --

