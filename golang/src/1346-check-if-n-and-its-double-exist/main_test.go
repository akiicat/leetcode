package main
import "testing"
import . "main/pkg/testing_helper"

func TestCheckIfExist(_t *testing.T) {
  i, o := []int{10,2,5,3}, true
  T(_t, S(i), S(checkIfExist(i)), S(o))

  i, o = []int{7,1,14,11}, true
  T(_t, S(i), S(checkIfExist(i)), S(o))

  i, o = []int{3,1,7,11}, false
  T(_t, S(i), S(checkIfExist(i)), S(o))
}

