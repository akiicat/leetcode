package main

// T: O(n)
// M: O(n)
// -- start --

func sortedSquares(A []int) []int {
  min_index := 0

  B := make([]int, len(A))

  for i, num := range A {
    A[i] = num * num
    if A[i] < A[min_index] {
      min_index = i
    }
  }

  for i, l, r := 0, min_index, min_index + 1; l >= 0 || r < len(A); i++ {

    if r >= len(A) || l >= 0 && A[l] < A[r] {
      B[i] = A[l]
      l--
    } else {
      B[i] = A[r]
      r++
    }

  }

  return B
}

func sortedSquaresReverse(A []int) []int {
  B := make([]int,len(A))

  last := len(A) - 1

  for i, start, end := last, 0, last; start <= end; i-- {
    // abs
    if A[start] < 0 {
      A[start] = 0 - A[start]
    }
    if A[end] < 0 {
      A[end] = 0 - A[end]
    }

    if A[start] > A[end] {
      B[i] = A[start] * A[start]
      start++
    } else {
      B[i] = A[end] * A[end]
      end--
    }
  }

  return B;
}

// -- end --

