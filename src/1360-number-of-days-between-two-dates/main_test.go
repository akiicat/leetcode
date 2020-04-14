package main
import "testing"
import . "main/pkg/testing_helper"

func TestDaysBetweenDates(_t *testing.T) {
  i1, i2, o := "2019-06-29", "2019-06-30", 1
  T(_t, S(i1, i2), S(daysBetweenDates(i1, i2)), S(o))

  i1, i2, o = "2020-01-15", "2019-12-31", 15
  T(_t, S(i1, i2), S(daysBetweenDates(i1, i2)), S(o))

  i1, i2, o = "2019-06-30", "2019-06-29", 1
  T(_t, S(i1, i2), S(daysBetweenDates(i1, i2)), S(o))
}

func TestDaysBetweenDatesStrconv(_t *testing.T) {
  i1, i2, o := "2019-06-29", "2019-06-30", 1
  T(_t, S(i1, i2), S(daysBetweenDatesStrconv(i1, i2)), S(o))

  i1, i2, o = "2020-01-15", "2019-12-31", 15
  T(_t, S(i1, i2), S(daysBetweenDatesStrconv(i1, i2)), S(o))

  i1, i2, o = "2019-06-30", "2019-06-29", 1
  T(_t, S(i1, i2), S(daysBetweenDatesStrconv(i1, i2)), S(o))
}

