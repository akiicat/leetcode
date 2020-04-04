package main
import "testing"
import . "main/pkg/testing_helper"

func TestGetNoZeroIntegers(t *testing.T) {
  i, o := 2, []int{2,2}
  T(t, S(i), S(SumAndLen(getNoZeroIntegers(i))), S(o))

  i, o = 11, []int{11,2}
  T(t, S(i), S(SumAndLen(getNoZeroIntegers(i))), S(o))

  i, o = 10000, []int{10000,2}
  T(t, S(i), S(SumAndLen(getNoZeroIntegers(i))), S(o))

  i, o = 69, []int{69,2}
  T(t, S(i), S(SumAndLen(getNoZeroIntegers(i))), S(o))

  i, o = 1010, []int{1010,2}
  T(t, S(i), S(SumAndLen(getNoZeroIntegers(i))), S(o))
}

func SumAndLen(arr []int) []int {
  sum := 0
  for _, v := range arr {
    zero := false

    for i := v; i != 0; i /= 10 {
      if i % 10 == 0 {
        zero = true
      }
    }

    if !zero {
      sum += v
    }
  }
  return []int{sum, len(arr)}
}
