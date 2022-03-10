package main

// T: O(n)
// M: O(n)
// -- start --

// T: O(n)
// M: O(n)

type indices struct {
  s, e, c int     // start, end, count
}

func findShortestSubArray(nums []int) int {
  m := make(map[int]*indices)

  max := 0 // max count
  for i, v := range nums {
    _, ok := m[v]

    if !ok {
      m[v] = &indices{s: i, e: i, c: 1}
      max = Max(max, 1)
      continue
    }

    m[v].e = i
    m[v].c++
    max = Max(max, m[v].c)
  }

  min := len(nums) // min length
  for _, v := range m {
    if v.c == max {
      min = Min(min, v.e - v.s + 1)
    }
  }

  return min
}

// T: O(n)
// M: O(n)
func findShortestSubArrayMap(nums []int) int {
  m := make(map[int]int)
  l := make(map[int]int)
  r := make(map[int]int)

  max := 0
  for i, v := range nums {
    m[v]++
    max = Max(max, m[v])

    _, ok := l[v]
    if !ok {
      l[v] = i
    }

    r[v] = i
  }

  length := len(nums)
  for i, v := range m {
    if v != max {
      continue
    }

    length = Min(length, r[i] - l[i] + 1)
  }

  return length
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func Min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

