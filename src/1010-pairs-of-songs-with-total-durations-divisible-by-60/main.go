package main

// T: O(n)
// M: O(1) for 60
// -- start --

func numPairsDivisibleBy60(time []int) int {
  m := [60]int{}
  sum := 0

  for _, v := range time {
    i := v % 60
    r := (60 - i) % 60
    sum += m[r]
    m[i]++
  }

  return sum
}

// -- end --

