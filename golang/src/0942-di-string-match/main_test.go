package main
import "testing"
import . "main/pkg/testing_helper"

func TestDiStringMatch(t *testing.T) {
  i, o := "IDID", []int{0,4,1,3,2}
  T(t, S(i), S(diStringMatch(i)), S(o))

  i, o = "III", []int{0,1,2,3}
  T(t, S(i), S(diStringMatch(i)), S(o))

  i, o = "DDI", []int{3,2,0,1}
  T(t, S(i), S(diStringMatch(i)), S(o))
}

