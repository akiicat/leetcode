package main
import "testing"
import . "main/pkg/testing_helper"

func TestFourSum(_t *testing.T) {
  i, t, o := []int{1,0,-1,0,-2,2}, 0, [][]int{
    []int{-1,0,0,1},
    []int{-2,-1,1,2},
    []int{-2,0,0,2},
  }
  T(_t, S(i, t), S(valid(fourSum(i, t), o, t)), S(o))

  i, t, o = []int{-1,0,1,2,-1,-4}, -1, [][]int{
    []int{-4,0,1,2},
    []int{-1,-1,0,1},
  }
  T(_t, S(i, t), S(valid(fourSum(i, t), o, t)), S(o))

  i, t, o = []int{-3,-2,-1,0,0,1,2,3}, 0, [][]int{
    []int{-3,-2,2,3},
    []int{-3,-1,1,3},
    []int{-3,0,0,3},
    []int{-3,0,1,2},
    []int{-2,-1,0,3},
    []int{-2,-1,1,2},
    []int{-2,0,0,2},
    []int{-1,0,0,1},
  }
  T(_t, S(i, t), S(valid(fourSum(i, t), o, t)), S(o))

  i, t, o = []int{-1,0,-5,-2,-2,-4,0,1,-2}, -9, [][]int{
    []int{-5,-4,-1,1},
    []int{-5,-4,0,0},
    []int{-5,-2,-2,0},
    []int{-4,-2,-2,-1},
  }
  T(_t, S(i, t), S(valid(fourSum(i, t), o, t)), S(o))
}

func valid(r, o [][]int, t int) [][]int {
  if len(r) != len(o) {
    return r
  }

  for _, v := range r {
    if len(v) != 4 || sum(v) != t {
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


