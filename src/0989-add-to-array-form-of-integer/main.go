package main

// T: O(max(n, log(k))) n is the length of A
// M: O(max(n, log(k)))
// -- start --

func addToArrayForm(A []int, K int) []int {
  for i := len(A) - 1; i >= 0 || K > 0; i-- {
    if i >= 0 {
      K += A[i]
      A[i] = K % 10
    } else {
      A = append([]int{K % 10}, A...)
    }

    K /= 10
  }

  return A
}

// -- end --

