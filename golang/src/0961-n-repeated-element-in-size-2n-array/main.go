package main

// T: O(n)
// M: O(1)
// -- start --

// Pigeonhole Principle
// T: O(n)
// M: O(1)
func repeatedNTimes(A []int) int {
  for k := 0; k < 3; k++ {
    for i, v := range A[k+1:] {
      if v == A[i] {
        return v
      }
    }
  }
  return 0
}

// T: O(n)
// M: O(n)
func repeatedNTimesArray(A []int) int {
  m := [10000]int{}
  for _, v := range A {
    if m[v] != 0 {
      return v
    }
    m[v]++
  }
  return 0
}

// T: O(n)
// M: O(n)
func repeatedNTimesHash(A []int) int {
  m := make(map[int]bool)
  for _, v := range A {
    if m[v] {
      return v
    }
    m[v] = true
  }
  return 0
}

// -- end --

