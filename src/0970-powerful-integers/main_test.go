package main
import "testing"
import . "main/pkg/testing_helper"

func TestPowerfulIntegers(t *testing.T) {
  x, y, bound, o := 2, 3, 10, []int{2,3,4,5,7,9,10}
  T(t, S(x, y, bound), S(Sort(powerfulIntegers(x, y, bound))), S(Sort(o)))

  x, y, bound, o = 3, 5, 15, []int{2,4,6,8,10,14}
  T(t, S(x, y, bound), S(Sort(powerfulIntegers(x, y, bound))), S(Sort(o)))

  x, y, bound, o = 2, 1, 10, []int{2,3,5,9}
  T(t, S(x, y, bound), S(Sort(powerfulIntegers(x, y, bound))), S(Sort(o)))
}

func TestPowerfulIntegersPower(t *testing.T) {
  x, y, bound, o := 2, 3, 10, []int{2,3,4,5,7,9,10}
  T(t, S(x, y, bound), S(Sort(powerfulIntegersPower(x, y, bound))), S(Sort(o)))

  x, y, bound, o = 3, 5, 15, []int{2,4,6,8,10,14}
  T(t, S(x, y, bound), S(Sort(powerfulIntegersPower(x, y, bound))), S(Sort(o)))

  x, y, bound, o = 2, 1, 10, []int{2,3,5,9}
  T(t, S(x, y, bound), S(Sort(powerfulIntegersPower(x, y, bound))), S(Sort(o)))
}

