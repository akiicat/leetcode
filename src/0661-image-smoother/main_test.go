package main
import "testing"
import . "main/pkg/testing_helper"

func TestImageSmoother(_t *testing.T) {
  i, o := [][]int{[]int{1,1,1},[]int{1,0,1},[]int{1,1,1}}, [][]int{[]int{0,0,0},[]int{0,0,0},[]int{0,0,0}}
  T(_t, S(i), S(imageSmoother(i)), S(o))

  i, o = [][]int{[]int{1}}, [][]int{[]int{1}}
  T(_t, S(i), S(imageSmoother(i)), S(o))

  i, o = [][]int{[]int{2,3}}, [][]int{[]int{2,2}}
  T(_t, S(i), S(imageSmoother(i)), S(o))

  i, o = [][]int{[]int{2},[]int{3}}, [][]int{[]int{2},[]int{2}}
  T(_t, S(i), S(imageSmoother(i)), S(o))
}
