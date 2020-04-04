package main
import "testing"
import . "main/pkg/testing_helper"

func TestSumZero(t *testing.T) {
  i, o := 5, 0
  r := sumZero(i)
  T(t, S(i), S(len(r) + sum(r)), S(i + o))

  i, o = 4, 0
  r = sumZero(i)
  T(t, S(i), S(len(r) + sum(r)), S(i + o))

  i, o = 3, 0
  r = sumZero(i)
  T(t, S(i), S(len(r) + sum(r)), S(i + o))

  i, o = 1, 0
  r = sumZero(i)
  T(t, S(i), S(len(r) + sum(r)), S(i + o))
}

func sum(a []int) int {
  sum := 0
  for _, v := range a {
    sum += v
  }
  return sum
}
