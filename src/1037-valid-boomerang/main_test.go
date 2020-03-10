package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsBoomerang(_t *testing.T) {
  i, o := [][]int{[]int{1,1},[]int{2,3},[]int{3,2}}, true
  T(_t, S(i), S(isBoomerang(i)), S(o))

  i, o = [][]int{[]int{1,1},[]int{2,2},[]int{3,3}}, false
  T(_t, S(i), S(isBoomerang(i)), S(o))

  i, o = [][]int{[]int{0,1},[]int{0,1},[]int{2,1}}, false
  T(_t, S(i), S(isBoomerang(i)), S(o))

  i, o = [][]int{[]int{0,1},[]int{1,0},[]int{0,1}}, false
  T(_t, S(i), S(isBoomerang(i)), S(o))
}

