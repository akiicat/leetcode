package main
import "testing"
import . "main/pkg/testing_helper"

func TestThirdMax(_t *testing.T) {
  i, o := []int{3,2,1}, 1
  T(_t, S(i), S(thirdMax(i)), S(o))

  i, o = []int{1,2}, 2
  T(_t, S(i), S(thirdMax(i)), S(o))

  i, o = []int{2,2,3,1}, 1
  T(_t, S(i), S(thirdMax(i)), S(o))

  i, o = []int{5,2,2}, 5
  T(_t, S(i), S(thirdMax(i)), S(o))

  i, o = []int{-2147483648,1,1}, 1
  T(_t, S(i), S(thirdMax(i)), S(o))

  i, o = []int{1,2,-2147483648}, -2147483648
  T(_t, S(i), S(thirdMax(i)), S(o))

  i, o = []int{1,-2147483648,2}, -2147483648
  T(_t, S(i), S(thirdMax(i)), S(o))
}

func TestThirdMaxMap(_t *testing.T) {
  i, o := []int{3,2,1}, 1
  T(_t, S(i), S(thirdMaxMap(i)), S(o))

  i, o = []int{1,2}, 2
  T(_t, S(i), S(thirdMaxMap(i)), S(o))

  i, o = []int{2,2,3,1}, 1
  T(_t, S(i), S(thirdMaxMap(i)), S(o))

  i, o = []int{5,2,2}, 5
  T(_t, S(i), S(thirdMaxMap(i)), S(o))

  i, o = []int{-2147483648,1,1}, 1
  T(_t, S(i), S(thirdMaxMap(i)), S(o))

  i, o = []int{1,2,-2147483648}, -2147483648
  T(_t, S(i), S(thirdMaxMap(i)), S(o))

  i, o = []int{1,-2147483648,2}, -2147483648
  T(_t, S(i), S(thirdMaxMap(i)), S(o))
}
