package main
import "testing"
import . "main/pkg/testing_helper"

func TestGetSum(_t *testing.T) {
  r, m, o := "a", "b", false
  T(_t, S(r, m), S(canConstruct(r, m)), S(o))

  r, m, o = "aa", "ab", false
  T(_t, S(r, m), S(canConstruct(r, m)), S(o))

  r, m, o = "aa", "aab", true
  T(_t, S(r, m), S(canConstruct(r, m)), S(o))

  r, m, o = "aab", "aa", false
  T(_t, S(r, m), S(canConstruct(r, m)), S(o))

  r, m, o = "bg", "abge", true
  T(_t, S(r, m), S(canConstruct(r, m)), S(o))

  r, m, o = "", "", true
  T(_t, S(r, m), S(canConstruct(r, m)), S(o))
}
