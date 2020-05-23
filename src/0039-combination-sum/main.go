package main

// T: O(n ^ k) n is the size of candidates, k is target
// M: O(k)
// -- start --

func combinationSum(candidates []int, target int) [][]int {
  res := [][]int{}

  R(&res, []int{}, 0, candidates, target)

  return res
}

func R(res *[][]int, arr []int, sum int, candidates []int, target int) {
  if sum > target {
    return
  }

  if sum == target {
    a := make([]int, len(arr))
    copy(a, arr)
    *res = append(*res, a)
  }

  for _, v := range candidates {
    if len(arr) == 0 || v >= arr[len(arr)-1] {
      R(res, append(arr, v), sum + v, candidates, target)
    }
  }
}

// -- end --

