package main

// T: O(n) n is the length of the string
// M: O(1)
// -- start --

// One Pass
func isMonotonic(A []int) bool {
  inc, dec := true, true

  for i := 0; i < len(A) - 1; i++ {
    if A[i+1] > A[i] {
      dec = false
    }

    if A[i+1] < A[i] {
      inc = false
    }
  }

  return inc || dec
}

// -- end --

