package main
import "testing"
import . "main/pkg/testing_helper"

func TestDayOfTheWeek(_t *testing.T) {
  d, m, y, o := 31, 8, 2019, "Saturday"
  T(_t, S(d, m, y), S(dayOfTheWeek(d, m, y)), S(o))

  d, m, y, o = 18, 7, 1999, "Sunday"
  T(_t, S(d, m, y), S(dayOfTheWeek(d, m, y)), S(o))

  d, m, y, o = 15, 8, 1993, "Sunday"
  T(_t, S(d, m, y), S(dayOfTheWeek(d, m, y)), S(o))
}

func TestDayOfTheWeekTime(_t *testing.T) {
  d, m, y, o := 31, 8, 2019, "Saturday"
  T(_t, S(d, m, y), S(dayOfTheWeekTime(d, m, y)), S(o))

  d, m, y, o = 18, 7, 1999, "Sunday"
  T(_t, S(d, m, y), S(dayOfTheWeekTime(d, m, y)), S(o))

  d, m, y, o = 15, 8, 1993, "Sunday"
  T(_t, S(d, m, y), S(dayOfTheWeekTime(d, m, y)), S(o))
}

