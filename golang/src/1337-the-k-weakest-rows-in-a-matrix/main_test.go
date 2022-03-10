package main
import "testing"
import . "main/pkg/testing_helper"

func TestKWeakestRows(_t *testing.T) {
  i, k, o := [][]int{
    []int{1,1,0,0,0},
    []int{1,1,1,1,0},
    []int{1,0,0,0,0},
    []int{1,1,0,0,0},
    []int{1,1,1,1,1},
  }, 3, []int{2,0,3}
  T(_t, S(i, k), S(kWeakestRows(i, k)), S(o))

  i, k, o = [][]int{
    []int{1,0,0,0},
    []int{1,1,1,1},
    []int{1,0,0,0},
    []int{1,0,0,0},
  }, 2, []int{0,2}
  T(_t, S(i, k), S(kWeakestRows(i, k)), S(o))

  i, k, o = [][]int{
    []int{1,1,0},
    []int{1,1,0},
    []int{1,1,1},
    []int{1,1,1},
    []int{0,0,0},
    []int{1,1,1},
    []int{1,0,0},
  }, 6, []int{4,6,0,1,2,3}
  T(_t, S(i, k), S(kWeakestRows(i, k)), S(o))
}

