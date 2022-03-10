package main
import "testing"
import . "main/pkg/testing_helper"

func TestSelfDividingNumbers(_t *testing.T) {
  l, r, o := 1, 22, []int{1,2,3,4,5,6,7,8,9,11,12,15,22}
  T(_t, S(l, r), S(selfDividingNumbers(l, r)), S(o))
}

