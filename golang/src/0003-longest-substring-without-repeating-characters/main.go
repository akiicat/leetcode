package main

// T: O(n)
// M: O(1) for 128
// -- start --

func lengthOfLongestSubstring(s string) int {
  m := [128]int{}

  l, score, max := 0, 0, 0
  for _, v := range s {
    m[int(v)]++
    score++

    for m[int(v)] > 1 {
      m[int(s[l])]--
      score--
      l++
    }

    if score > max {
      max = score
    }
  }

  return max
}

// -- end --

