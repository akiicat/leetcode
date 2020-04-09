package main

import "sort"

// T: O(n * log(n))
// M: O(n)
// -- start --

func kWeakestRows(mat [][]int, k int) []int {

  m := make([][2]int, len(mat))

  for i, row := range mat {
    m[i] = [2]int{i, Sum(row)}
  }

  sort.Slice(m, func(i, j int) bool {
    if m[i][1] == m[j][1] {
      return m[i][0] < m[j][0]
    }
    return m[i][1] < m[j][1]
  })

  rtn := make([]int, k)
  for i := 0; i < k; i++ {
    rtn[i] = m[i][0]
  }

  return rtn
}

func Sum(a []int) int {
  sum := 0
  for _, v := range a {
    sum += v
  }
  return sum
}

// -- end --

