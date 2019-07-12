package main
import "fmt"

func main() {
  a, k, o := []int{1, 2, 0, 0}, 34, []int{1, 2, 3, 4}
  fmt.Printf("Input:  %v, %d\n", a, k)
  fmt.Printf("Output: %v\nExpect: %v\n", addToArrayForm(a, k), o)

  a, k, o = []int{2, 7, 4}, 181, []int{4, 5, 5}
  fmt.Printf("Input:  %v, %d\n", a, k)
  fmt.Printf("Output: %v\nExpect: %v\n", addToArrayForm(a, k), o)

  a, k, o = []int{2, 1, 5}, 806, []int{1, 0, 2, 1}
  fmt.Printf("Input:  %v, %d\n", a, k)
  fmt.Printf("Output: %v\nExpect: %v\n", addToArrayForm(a, k), o)

  a, k, o = []int{9,9,9,9,9,9,9,9,9,9}, 1, []int{1,0,0,0,0,0,0,0,0,0,0}
  fmt.Printf("Input:  %v, %d\n", a, k)
  fmt.Printf("Output: %v\nExpect: %v\n", addToArrayForm(a, k), o)

  a, k, o = []int{0}, 23, []int{2, 3}
  fmt.Printf("Input:  %v, %d\n", a, k)
  fmt.Printf("Output: %v\nExpect: %v\n", addToArrayForm(a, k), o)

}

// T: O(max(N, log(K))) N is the length of A
// M: O(max(N, log(K)))
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

