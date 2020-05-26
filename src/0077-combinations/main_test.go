package main
import "testing"
import . "main/pkg/testing_helper"

func TestCombine(_t *testing.T) {
  n, k, o := 4, 2, [][]int{
    []int{2,4},
    []int{3,4},
    []int{2,3},
    []int{1,2},
    []int{1,3},
    []int{1,4},
  }
  T(_t, S(n, k), S(Sort(combine(n, k))), S(Sort(o)))
}

