package main
import "sort"

// T: O(n * log(n))
// M: O(n)
// -- start --

func heightChecker(heights []int) int {
  m := make([]int, len(heights))
  copy(m, heights)
  sort.Ints(m)

  sum := 0
  for i, v := range m {
    if v != heights[i] {
      sum++
    }
  }

  return sum
}

// -- end --
