package main

// T: O(n * log(n) ** 2)
// M: O(n)
// -- start --

func hasGroupsSizeX(deck []int) bool {
  m := make(map[int]int)
  for _, v := range deck {
    m[v]++
  }

  t := -1
  for _, v := range m {
    if t == -1 {
      t = v
    } else {
      t = gcd(t, v)
    }
  }

  return t >= 2
}

func gcd(a, b int) int {
  for b != 0 {
    a, b = b, a % b
  }
  return a
}

// -- end --
