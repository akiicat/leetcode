package main
import "testing"
import . "main/pkg/testing_helper"

func TestNumPrimeArrangements(_t *testing.T) {
  i, o := 2, 1
  T(_t, S(i), S(numPrimeArrangements(i)), S(o))

  i, o = 5, 12
  T(_t, S(i), S(numPrimeArrangements(i)), S(o))

  i, o = 10, 17280
  T(_t, S(i), S(numPrimeArrangements(i)), S(o))

  i, o = 11, 86400
  T(_t, S(i), S(numPrimeArrangements(i)), S(o))

  i, o = 100, 682289015
  T(_t, S(i), S(numPrimeArrangements(i)), S(o))
}

