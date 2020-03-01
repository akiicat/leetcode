package main
import "sort"

// T: O(n + m)
// M: O(n)
// -- start --

// T: O(n + m)
// M: O(n)
func fairCandySwap(A []int, B []int) []int {
  m := make(map[int]bool, len(B))
  g := 0
  for _, v := range A {
    g -= v
  }
  for _, v := range B {
    g += v
    m[v] = true
  }
  g >>= 1

  for _, v := range A {
    if m[v+g] {
      return []int{v, v+g}
    }
  }

  return []int{0,0}
}

// T: O(max(n*log(n), m*log(m)))
// M: O(1)
func fairCandySwapSort(A []int, B []int) []int {
  sort.Ints(A)
  sort.Ints(B)

  g := Diff(A, B)

  for i, j := 0, 0; i < len(A) && j < len(B); {
    if B[j] - A[i] > g {
      i++
    } else if B[j] - A[i] < g {
      j++
    } else {
      return []int{A[i], B[j]}
    }
  }

  return []int{0,0}
}

func Diff(A []int, B []int) int {
  g := 0
  for _, v := range A {
    g -= v
  }
  for _, v := range B {
    g += v
  }
  return g >> 1
}

// -- end --

