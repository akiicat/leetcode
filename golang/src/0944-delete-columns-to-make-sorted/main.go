package main

// T: O(n) n is the total element of A
// M: O(1)
// -- start --

func minDeletionSize(A []string) int {
  if len(A) == 0 {
    return 0
  }

  count := 0

  for i := 0; i < len(A[0]); i++ {
    for k := 0; k < len(A)-1; k++ {
      if A[k][i] > A[k+1][i] {
        count++
        break
      }
    }
  }

  return count
}

// -- end --

