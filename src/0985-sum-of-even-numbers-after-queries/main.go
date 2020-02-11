package main

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

