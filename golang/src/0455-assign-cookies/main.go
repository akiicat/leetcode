package main
import "sort"

// T: O(n*log(n)) for sort
// M: O(1)
// -- start --

func findContentChildren(g []int, s []int) int {
  gn, sn := len(g), len(s)
  if gn == 0 || sn == 0 {
    return 0
  }

  sort.Ints(g)
  sort.Ints(s)

  i, j, count := 0, 0, 0
  for i < gn && j < sn {
    if g[i] <= s[j] {
      count++
      i++
    }
    j++
  }

  return count
}

// -- end --

