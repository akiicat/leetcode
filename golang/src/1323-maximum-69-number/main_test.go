package main
import "testing"
import . "main/pkg/testing_helper"

func TestMaximum69Number(t *testing.T) {
  i, o := 9669, 9969
  T(t, S(i), S(maximum69Number(i)), S(o))

  i, o = 9996, 9999
  T(t, S(i), S(maximum69Number(i)), S(o))

  i, o = 9999, 9999
  T(t, S(i), S(maximum69Number(i)), S(o))
}

func TestMaximum69NumberStrconv(t *testing.T) {
  i, o := 9669, 9969
  T(t, S(i), S(maximum69NumberStrconv(i)), S(o))

  i, o = 9996, 9999
  T(t, S(i), S(maximum69NumberStrconv(i)), S(o))

  i, o = 9999, 9999
  T(t, S(i), S(maximum69NumberStrconv(i)), S(o))
}

