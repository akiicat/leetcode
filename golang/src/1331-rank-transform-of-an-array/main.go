package main

import "sort"

// T: O(n * log(n))
// M: O(n)
// -- start --

func arrayRankTransform(arr []int) []int {
  if len(arr) == 0 {
    return nil
  }

  m := map[int]int{}
  a := make([]int, len(arr))
  copy(a, arr)

  sort.Ints(a)

  r, c := 1, a[0]
  for _, v := range a {
    if c < v {
      r, c = r+1, v
    }

    m[v] = r
  }

  for i, v := range arr {
    arr[i] = m[v]
  }

  return arr
}

// -- end --

