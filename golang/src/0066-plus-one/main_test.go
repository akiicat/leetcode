package main
import "testing"
import . "main/pkg/testing_helper"

func TestPlusOne(_t *testing.T) {
  i, o := []int{1,2,3}, []int{1,2,4}
  T(_t, S(i), S(plusOne(i)), S(o))

  i, o = []int{4,3,2,1}, []int{4,3,2,2}
  T(_t, S(i), S(plusOne(i)), S(o))

  i, o = []int{4,3,2,9}, []int{4,3,3,0}
  T(_t, S(i), S(plusOne(i)), S(o))
}
