package main

// T: O(n)
// M: O(1)
// -- start --

func validMountainArray(A []int) bool {
  n := len(A)

  if n < 3 || A[0] >= A[1] || A[n-2] <= A[n-1] {
    return false
  }

  i := 0
  for i < n-1 && A[i] < A[i+1] {
    i++
  }
  for i < n-1 && A[i] > A[i+1] {
    i++
  }

  return i+1 == n
}

// -- end --

