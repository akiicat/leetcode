package main

// T: O(n)
// M: O(1) for 100
// -- start --

func numEquivDominoPairs(dominoes [][]int) int {
  m := [100]int{}

  for _, v := range dominoes {
    if v[0] > v[1] {
      v[0], v[1] = v[1], v[0]
    }
    m[10 * v[0] + v[1]]++
  }

  sum := 0
  for _, v := range m {
    if v != 1 {
      sum += (v * (v-1)) / 2
    }
  }

  return sum
}

// -- end --

