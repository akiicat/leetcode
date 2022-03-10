package main
import "testing"
import . "main/pkg/testing_helper"

func TestMoveZeroes(_t *testing.T) {
  i, o := []int{0,1,0,3,12}, []int{1,3,12,0,0}
  _i := S(i)
  moveZeroes(i)
  T(_t, _i, S(i), S(o))

  i, o = []int{2,1}, []int{2,1}
  _i = S(i)
  moveZeroes(i)
  T(_t, _i, S(i), S(o))
}
