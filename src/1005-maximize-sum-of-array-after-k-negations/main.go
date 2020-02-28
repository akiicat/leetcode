package main
import "sort"

// T: O(n) n ~ n**2, n on average
// M: O(1)
// -- start --

func largestSumAfterKNegations(A []int, K int) int {
  for K > 0 {
    for i := 1; i < len(A); i++ {
      if A[0] > A[i] {
        A[0], A[i] = A[i], A[0]
      }
    }

    A[0], K = -A[0], K-1
  }

  sum := 0
  for _, v := range A {
    sum += v
  }
  return sum
}


// T: O(n * log(n))
// M: O(1)
func largestSumAfterKNegationsSort(A []int, K int) int {
  sort.Ints(A)
  sum := 0
  for i := 0; i < len(A); i++ {
    if A[i] < 0 && K > 0 {
      A[i], K = -A[i], K-1
    }

    sum += A[i]

    if A[i] < A[0] {
      A[0], A[i] = A[i], A[0]
    }
  }

  if K & 0x1 == 1 {
    sum -= 2*A[0]
  }

  return sum
}

// -- end --

