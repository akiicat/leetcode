package main

// T: O(n)
// M: O(1)
// -- start --

func isPalindrome(x int) bool {
  if x < 0 {
    return false
  }

  y := 0
  tmp := x
  for tmp != 0 {
    y = 10 * y + tmp % 10
    tmp /= 10
  }

  return x == y
}

// -- end --

