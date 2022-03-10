package main

// leetcode 922.
// T: O(n)
// M: O(1)
// -- start --

func sortArrayByParity(A []int) []int {
  i := 0
  for j, v := range A {
    if v & 0x1 == 0 {
      A[i], A[j] = v, A[i]
      i++
    }
  }

  return A
}

// -- end --

