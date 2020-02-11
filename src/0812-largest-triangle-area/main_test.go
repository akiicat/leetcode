package main
import "testing"
import . "main/pkg/testing_helper"

func TestLargestTriangleArea(_t *testing.T) {
  i, o := [][]int{[]int{0,0},[]int{0,1},[]int{1,0},[]int{0,2},[]int{2,0}}, 2.0
  T(_t, S(i), S(largestTriangleArea(i)), S(o))
}

