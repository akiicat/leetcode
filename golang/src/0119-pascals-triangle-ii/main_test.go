package main
import "testing"
import . "main/pkg/testing_helper"

func TestGetRow(_t *testing.T) {
  i, o := 3, []int{1,3,3,1}
  T(_t, S(i), S(getRow(i)), S(o))

  i, o = 0, []int{1}
  T(_t, S(i), S(getRow(i)), S(o))

  i, o = 4, []int{1,4,6,4,1}
  T(_t, S(i), S(getRow(i)), S(o))
}
