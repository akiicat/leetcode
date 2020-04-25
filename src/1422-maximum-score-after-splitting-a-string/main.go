package main

// T: O(n)
// M: O(1)
// -- start --

func maxScore(s string) int {
  l, r := 0, 0
  for _, v := range s[1:len(s)-1] {
    if v == '1' {
      r++
    }
  }

  max := 0

  for _, v := range s[1:] {
    if l + r > max {
      max = l + r
    }

    if v == '0' {
      l++
    } else { // v == '1'
      r--
    }
  }

  if s[0] == '0' {
    max++
  }
  if s[len(s)-1] == '1' {
    max++
  }

  return max
}

// -- end --

