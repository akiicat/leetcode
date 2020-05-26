package main

// T: O(k * C(n,k))
// M: O(k)
// -- start --

func combine(n int, k int) [][]int {
  res := [][]int{}
  R(&res, []int{}, 1, 0, n, k)
  return res
}

func R(res *[][]int, arr []int, cur, count, n, k int) {
  if count == k {
    a := make([]int, len(arr))
    copy(a, arr)
    *res = append(*res, a)
    return
  }

  for i := cur; i <= n; i++ {
    R(res, append(arr, i), i+1, count+1, n, k)
  }
}

// -- end --

