package main
import "testing"
import . "main/pkg/testing_helper"

func TestPivotIndex(_t *testing.T) {
  i, o := []int{1,7,3,6,5,6}, 3
  T(_t, S(i), S(pivotIndex(i)), S(o))

  i, o = []int{1,2,3}, -1
  T(_t, S(i), S(pivotIndex(i)), S(o))
}

