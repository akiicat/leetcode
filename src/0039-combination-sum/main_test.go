package main
import "testing"
import . "main/pkg/testing_helper"

func TestCombinationSum(_t *testing.T) {
  c, t, o := []int{2,3,6,7}, 7, [][]int{[]int{7},[]int{2,2,3}}
  T(_t, S(c,t), S(Sort(combinationSum(c,t))), S(Sort(o)))

  c, t, o = []int{2,3,5}, 8, [][]int{[]int{2,2,2,2},[]int{2,3,3},[]int{3,5}}
  T(_t, S(c,t), S(Sort(combinationSum(c,t))), S(Sort(o)))
}

