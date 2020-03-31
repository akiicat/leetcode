package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindNumbers(t *testing.T) {
  i, o := []int{12,345,2,6,7896}, 2
  T(t, S(i), S(findNumbers(i)), S(o))

  i, o = []int{555,901,482,1771}, 1
  T(t, S(i), S(findNumbers(i)), S(o))
}

