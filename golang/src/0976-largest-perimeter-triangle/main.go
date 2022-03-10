package main
import "sort"

// T: O(n*log(n)) n for the number of A
// M: O(1)
// -- start --

func largestPerimeter(A []int) int {
  sort.Sort(sort.Reverse(sort.IntSlice(A)))  // n*log(n)

  for i := 0; i < len(A)-2; i++ {
    if A[i] < A[i+1] + A[i+2] {
      return A[i] + A[i+1] + A[i+2]
    }
  }

  return 0
}

// -- end --

