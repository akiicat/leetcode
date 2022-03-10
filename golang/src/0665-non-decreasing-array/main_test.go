package main
import "testing"
import . "main/pkg/testing_helper"

func TestCheckPossibility(_t *testing.T) {
  i, o := []int{4,2,3}, true
  T(_t, S(i), S(checkPossibility(i)), S(o))

  i, o = []int{4,2,1}, false
  T(_t, S(i), S(checkPossibility(i)), S(o))

  i, o = []int{3,4,2,3}, false
  T(_t, S(i), S(checkPossibility(i)), S(o))

  i, o = []int{2,3,3,2,4}, true
  T(_t, S(i), S(checkPossibility(i)), S(o))
}
