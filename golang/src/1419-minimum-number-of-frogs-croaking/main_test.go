package main
import "testing"
import . "main/pkg/testing_helper"

func TestMinNumberOfFrogs(_t *testing.T) {
  i, o := "croakcroak", 1
  T(_t, S(i), S(minNumberOfFrogs(i)), S(o))

  i, o = "crcoakroak", 2
  T(_t, S(i), S(minNumberOfFrogs(i)), S(o))

  i, o = "ccrrooaakk", 2
  T(_t, S(i), S(minNumberOfFrogs(i)), S(o))

  i, o = "ccrrooaakkccrrooaakkcroak", 2
  T(_t, S(i), S(minNumberOfFrogs(i)), S(o))

  i, o = "crocakcroraoakk", 2
  T(_t, S(i), S(minNumberOfFrogs(i)), S(o))

  i, o = "ccrrooaakkccrrooaackkroak", 3
  T(_t, S(i), S(minNumberOfFrogs(i)), S(o))

  i, o = "croakcrook", -1
  T(_t, S(i), S(minNumberOfFrogs(i)), S(o))

  i, o = "croakcroa", -1
  T(_t, S(i), S(minNumberOfFrogs(i)), S(o))

  i, o = "aoocrrackk", -1
  T(_t, S(i), S(minNumberOfFrogs(i)), S(o))
}

