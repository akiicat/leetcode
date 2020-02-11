package main

// T: O(n)
// M: O(1)
// -- start --

func balancedStringSplit(s string) int {

  count, b := 0, 0

  for _, v := range s {
    switch v {
    case 'R': b++
    case 'L': b--
    }
    if b == 0 {
      count++
    }
  }

  return count
}

// -- end --

