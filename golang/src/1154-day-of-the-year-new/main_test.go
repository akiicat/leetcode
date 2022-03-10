package main
import "testing"
import . "main/pkg/testing_helper"

func TestDayOfYear(t *testing.T) {
  i, o := "2019-01-09", 9
  T(t, S(i), S(dayOfYear(i)), S(o))

  i, o = "2019-02-10", 41
  T(t, S(i), S(dayOfYear(i)), S(o))

  i, o = "2003-03-01", 60
  T(t, S(i), S(dayOfYear(i)), S(o))

  i, o = "2016-02-29", 60
  T(t, S(i), S(dayOfYear(i)), S(o))

  i, o = "2004-03-01", 61
  T(t, S(i), S(dayOfYear(i)), S(o))

  i, o = "1900-03-25", 84
  T(t, S(i), S(dayOfYear(i)), S(o))

  i, o = "2000-03-25", 85
  T(t, S(i), S(dayOfYear(i)), S(o))
}

