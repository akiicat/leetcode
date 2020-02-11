package main

// T: O(1)
// M: O(1)
// -- start --

func convertToTitle(n int) string {
  s := ""

  for n = n - 1; n >= 0; n = n / 26 - 1 {
    s = string(n % 26 + 65) + s
  }

  return s
}

// -- end --

