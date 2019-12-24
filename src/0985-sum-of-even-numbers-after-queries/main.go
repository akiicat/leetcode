package main
import "fmt"

func main() {
  i, queries, o := []int{1,2,3,4}, [][]int{[]int{1,0},[]int{-3,1},[]int{-4,0},[]int{2,3}}, []int{8,6,2,4}
  fmt.Printf("Input:  %v, %v, %v\n", i, queries, o)
  fmt.Printf("Output: %d\n", sumEvenAfterQueries(i, queries))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n + q)
// M: O(q)
// -- start --

func sumEvenAfterQueries(A []int, queries [][]int) []int {
  sum := 0
  for _, v := range A {
    if v & 0x1 == 0 {
      sum += v
    }
  }

  rtn := make([]int, len(queries))
  for i, q := range queries {
    add, vi := q[0], q[1]

    if A[vi] & 0x1 == 0 {
      sum -= A[vi]
    }

    A[vi] += add

    if A[vi] & 0x1 == 0 {
      sum += A[vi]
    }

    rtn[i] = sum
  }

  return rtn
}

// -- end --

