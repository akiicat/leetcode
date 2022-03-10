package main
import "testing"
import . "main/pkg/testing_helper"

func TestDuplicateZeros(_t *testing.T) {
  i, o := []int{1,0,2,3,0,4,5,0}, []int{1,0,0,2,3,0,0,4}
  _i := S(i)
  duplicateZeros(i)
  T(_t, _i, S(i), S(o))

  i, o = []int{1,2,3}, []int{1,2,3}
  _i = S(i)
  duplicateZeros(i)
  T(_t, _i, S(i), S(o))

  i, o = []int{0,0,0,0,0,0,0}, []int{0,0,0,0,0,0,0}
  _i = S(i)
  duplicateZeros(i)
  T(_t, _i, S(i), S(o))

  i, o = []int{8,4,5,0,0,0,0,7}, []int{8,4,5,0,0,0,0,0}
  _i = S(i)
  duplicateZeros(i)
  T(_t, _i, S(i), S(o))
}

func TestDuplicateZerosCopy(_t *testing.T) {
  i, o := []int{1,0,2,3,0,4,5,0}, []int{1,0,0,2,3,0,0,4}
  _i := S(i)
  duplicateZerosCopy(i)
  T(_t, _i, S(i), S(o))

  i, o = []int{1,2,3}, []int{1,2,3}
  _i = S(i)
  duplicateZerosCopy(i)
  T(_t, _i, S(i), S(o))

  i, o = []int{0,0,0,0,0,0,0}, []int{0,0,0,0,0,0,0}
  _i = S(i)
  duplicateZerosCopy(i)
  T(_t, _i, S(i), S(o))

  i, o = []int{8,4,5,0,0,0,0,7}, []int{8,4,5,0,0,0,0,0}
  _i = S(i)
  duplicateZerosCopy(i)
  T(_t, _i, S(i), S(o))
}

