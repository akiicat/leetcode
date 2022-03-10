package main
import "testing"
import . "main/pkg/testing_helper"

func TestMySqrtBinarySearch(_t *testing.T) {
  i, o := 0, 0
  T(_t, S(i), S(mySqrtBinarySearch(i)), S(o))

  i, o = 1, 1
  T(_t, S(i), S(mySqrtBinarySearch(i)), S(o))

  i, o = 4, 2
  T(_t, S(i), S(mySqrtBinarySearch(i)), S(o))

  i, o = 8, 2
  T(_t, S(i), S(mySqrtBinarySearch(i)), S(o))

  i, o = 9, 3
  T(_t, S(i), S(mySqrtBinarySearch(i)), S(o))

  i, o = 10, 3
  T(_t, S(i), S(mySqrtBinarySearch(i)), S(o))

  i, o = 36, 6
  T(_t, S(i), S(mySqrtBinarySearch(i)), S(o))
}

func TestMySqrt(_t *testing.T) {
  i, o := 0, 0
  T(_t, S(i), S(mySqrt(i)), S(o))

  i, o = 1, 1
  T(_t, S(i), S(mySqrt(i)), S(o))

  i, o = 4, 2
  T(_t, S(i), S(mySqrt(i)), S(o))

  i, o = 8, 2
  T(_t, S(i), S(mySqrt(i)), S(o))

  i, o = 9, 3
  T(_t, S(i), S(mySqrt(i)), S(o))

  i, o = 10, 3
  T(_t, S(i), S(mySqrt(i)), S(o))

  i, o = 36, 6
  T(_t, S(i), S(mySqrt(i)), S(o))
}

func TestMySqrtMath(_t *testing.T) {
  i, o := 0, 0
  T(_t, S(i), S(mySqrtMath(i)), S(o))

  i, o = 1, 1
  T(_t, S(i), S(mySqrtMath(i)), S(o))

  i, o = 4, 2
  T(_t, S(i), S(mySqrtMath(i)), S(o))

  i, o = 8, 2
  T(_t, S(i), S(mySqrtMath(i)), S(o))

  i, o = 9, 3
  T(_t, S(i), S(mySqrtMath(i)), S(o))

  i, o = 10, 3
  T(_t, S(i), S(mySqrtMath(i)), S(o))

  i, o = 36, 6
  T(_t, S(i), S(mySqrtMath(i)), S(o))
}
