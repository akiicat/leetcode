package main
import "sort"

// T: O(2 ^ n)
// M: O(n)
// -- start --

func combinationSum2(candidates []int, target int) [][]int {
  sort.Ints(candidates)
  res := [][]int{}
  R(&res, []int{}, 0, 0, candidates, target)
  return res
}

func R(res *[][]int, arr []int, cur, sum int, candidates []int, target int) {
  if sum == target {
    a := make([]int, len(arr))
    copy(a, arr)
    *res = append(*res, a)
    return
  }
  if sum > target {
    return
  }

  if cur < len(candidates) {
    R(res, append(arr, candidates[cur]), cur+1, sum+candidates[cur], candidates, target)
  }

  for cur+1 < len(candidates) && candidates[cur] == candidates[cur+1] {
    cur++
  }

  if cur < len(candidates) {
    R(res, arr, cur+1, sum, candidates, target)
  }
}

// -- end --

