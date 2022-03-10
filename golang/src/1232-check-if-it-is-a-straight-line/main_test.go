package main
import "testing"
import . "main/pkg/testing_helper"

func TestCheckStraightLine(t *testing.T) {
  i, o := [][]int{[]int{1,2},[]int{2,3},[]int{3,4},[]int{4,5},[]int{5,6},[]int{6,7}}, true
  T(t, S(i), S(checkStraightLine(i)), S(o))

  i, o = [][]int{[]int{1,1},[]int{2,2},[]int{3,4},[]int{4,5},[]int{5,6},[]int{7,7}}, false
  T(t, S(i), S(checkStraightLine(i)), S(o))
}

func TestCheckStraightLineDouble(t *testing.T) {
  i, o := [][]int{[]int{1,2},[]int{2,3},[]int{3,4},[]int{4,5},[]int{5,6},[]int{6,7}}, true
  T(t, S(i), S(checkStraightLineDouble(i)), S(o))

  i, o = [][]int{[]int{1,1},[]int{2,2},[]int{3,4},[]int{4,5},[]int{5,6},[]int{7,7}}, false
  T(t, S(i), S(checkStraightLineDouble(i)), S(o))
}

