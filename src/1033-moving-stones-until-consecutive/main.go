package main
import "sort"

// T: O(1)
// M: O(1)
// -- start --

func numMovesStones(a int, b int, c int) []int {
  m := []int{a, b, c}
  sort.Ints(m)

  if m[2] - m[0] == 2 {
    return []int{0,0}
  }

  if m[2] - m[1] <= 2 || m[1] - m[0] <= 2 {
    return []int{1, m[2]-m[0]-2}
  }

  return []int{2, m[2]-m[0]-2}
}

// -- end --

