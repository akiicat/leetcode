package main
import "testing"
import . "main/pkg/testing_helper"

func TestCombinationSum2(_t *testing.T) {
  c, t, o := []int{10,1,2,7,6,1,5}, 8, [][]int{[]int{1,7},[]int{1,2,5},[]int{2,6},[]int{1,1,6}}
  T(_t, S(c, t), S(Sort(combinationSum2(c, t))), S(Sort(o)))

  c, t, o = []int{2,5,2,1,2}, 5, [][]int{[]int{1,2,2},[]int{5}}
  T(_t, S(c, t), S(Sort(combinationSum2(c, t))), S(Sort(o)))
}

