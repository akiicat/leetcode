package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsToeplitzMatrix(_t *testing.T) {
  i, o := [][]int{[]int{1,2,3,4},[]int{5,1,2,3},[]int{9,5,1,2}}, true
  T(_t, S(i), S(isToeplitzMatrix(i)), S(o))

  i, o = [][]int{[]int{1,2},[]int{2,2}}, false
  T(_t, S(i), S(isToeplitzMatrix(i)), S(o))
}

