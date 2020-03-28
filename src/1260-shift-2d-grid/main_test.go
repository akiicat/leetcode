package main
import "testing"
import . "main/pkg/testing_helper"

func TestShiftGrid(_t *testing.T) {
  i, k, o := [][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}}, 1, [][]int{[]int{9,1,2},[]int{3,4,5},[]int{6,7,8}}
  T(_t, S(i, k), S(shiftGrid(i, k)), S(o))

  i, k, o = [][]int{[]int{3,8,1,9},[]int{19,7,2,5},[]int{4,6,11,10},[]int{12,0,21,13}}, 4, [][]int{[]int{12,0,21,13},[]int{3,8,1,9},[]int{19,7,2,5},[]int{4,6,11,10}}
  T(_t, S(i, k), S(shiftGrid(i, k)), S(o))

  i, k, o = [][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}}, 9, [][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}}
  T(_t, S(i, k), S(shiftGrid(i, k)), S(o))

  i, k, o = [][]int{[]int{1}}, 100, [][]int{[]int{1}}
  T(_t, S(i, k), S(shiftGrid(i, k)), S(o))
}

func TestShiftGridArray(_t *testing.T) {
  i, k, o := [][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}}, 1, [][]int{[]int{9,1,2},[]int{3,4,5},[]int{6,7,8}}
  T(_t, S(i, k), S(shiftGridArray(i, k)), S(o))

  i, k, o = [][]int{[]int{3,8,1,9},[]int{19,7,2,5},[]int{4,6,11,10},[]int{12,0,21,13}}, 4, [][]int{[]int{12,0,21,13},[]int{3,8,1,9},[]int{19,7,2,5},[]int{4,6,11,10}}
  T(_t, S(i, k), S(shiftGridArray(i, k)), S(o))

  i, k, o = [][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}}, 9, [][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}}
  T(_t, S(i, k), S(shiftGridArray(i, k)), S(o))

  i, k, o = [][]int{[]int{1}}, 100, [][]int{[]int{1}}
  T(_t, S(i, k), S(shiftGridArray(i, k)), S(o))
}

