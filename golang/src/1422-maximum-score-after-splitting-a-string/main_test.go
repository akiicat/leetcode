package main
import "testing"
import . "main/pkg/testing_helper"

func TestMaxScore(_t *testing.T) {
  i, o := "011101", 5
  T(_t, S(i), S(maxScore(i)), S(o))

  i, o = "00111", 5
  T(_t, S(i), S(maxScore(i)), S(o))

  i, o = "1111", 3
  T(_t, S(i), S(maxScore(i)), S(o))

  i, o = "01001", 4
  T(_t, S(i), S(maxScore(i)), S(o))
}

