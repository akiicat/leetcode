package main
import "testing"
import . "main/pkg/testing_helper"

func TestConstructRectangle(_t *testing.T) {
  i, o := 4, []int{2,2}
  T(_t, S(i), S(constructRectangle(i)), S(o))

  i, o = 5, []int{5,1}
  T(_t, S(i), S(constructRectangle(i)), S(o))
}

