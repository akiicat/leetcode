package main
import "testing"
import . "main/pkg/testing_helper"

func TestGcdOfStrings(_t *testing.T) {
  i1, i2, o := "ABCABC", "ABC", "ABC"
  T(_t, S(i1, i2), S(gcdOfStrings(i1, i2)), S(o))

  i1, i2, o = "ABABAB", "ABAB", "AB"
  T(_t, S(i1, i2), S(gcdOfStrings(i1, i2)), S(o))

  i1, i2, o = "AAAAA", "AA", "A"
  T(_t, S(i1, i2), S(gcdOfStrings(i1, i2)), S(o))

  i1, i2, o = "AA", "AA", "AA"
  T(_t, S(i1, i2), S(gcdOfStrings(i1, i2)), S(o))

  i1, i2, o = "ACA", "AC", ""
  T(_t, S(i1, i2), S(gcdOfStrings(i1, i2)), S(o))

  i1, i2, o = "AA", "AC", ""
  T(_t, S(i1, i2), S(gcdOfStrings(i1, i2)), S(o))

  i1, i2, o = "LEET", "CODE", ""
  T(_t, S(i1, i2), S(gcdOfStrings(i1, i2)), S(o))

  i1, i2, o = "TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXX"
  T(_t, S(i1, i2), S(gcdOfStrings(i1, i2)), S(o))

}

