package main

// T: O(n)
// M: O(1) 26 chars
// -- start --

func maxNumberOfBalloons(text string) int {
  b := map[rune]int{
    'b': 1,
    'o': 2,
    'l': 2,
    'a': 1,
    'n': 1,
  }

  m := make(map[rune]int)

  for _, v := range text {
    m[v]++
  }

  min := 1<<31
  for k, v := range b {
    min = Min(min, m[k] / v)
  }

  return min
}

func Min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

