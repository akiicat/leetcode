package main

// T: O(n)
// M: O(1) for 26
// -- start --

func minNumberOfFrogs(croakOfFrogs string) int {
  c, r, o, a, k := 0, 0, 0, 0, 0
  res := 0

  for i := 0; i < len(croakOfFrogs); i++ {

    switch croakOfFrogs[i] {
    case 'c': c++
    case 'r': r++
    case 'o': o++
    case 'a': a++
    case 'k': k++
    }

    if c < r || r < o || o < a || a < k {
      return -1
    }

    res = Max(res, c-k)
  }

  if !(c == r && r == o && o == a && a == k) {
    return -1
  }

  return res
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

