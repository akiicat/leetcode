package main
import "testing"
import . "main/pkg/testing_helper"

func TestThreeSum(_t *testing.T) {
  i, o := []int{-1,0,1,2,-1,-4}, [][]int{[]int{-1,0,1},[]int{-1,-1,2}}
  T(_t, S(i), S(valid(threeSum(i), o)), S(o))

  i, o = []int{0,0,0,0}, [][]int{[]int{0,0,0}}
  T(_t, S(i), S(valid(threeSum(i), o)), S(o))

  i, o = []int{-2,0,1,1,2}, [][]int{[]int{-2,0,2},[]int{-2,1,1}}
  T(_t, S(i), S(valid(threeSum(i), o)), S(o))
}

func valid(r, o [][]int) [][]int {
  if len(r) != len(o) {
    return r
  }

  for _, v := range r {
    if len(v) != 3 || sum(v) != 0 {
      return r
    }
  }

  return o
}

func sum(a []int) int {
  s := 0
  for _, v := range a {
    s += v
  }
  return s
}


