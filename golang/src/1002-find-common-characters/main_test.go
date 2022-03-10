package main
import "testing"
import . "main/pkg/testing_helper"

func TestCommonChars(t *testing.T) {
  i, o := []string{"bella","label","roller"}, []string{"e","l","l"}
  T(t, S(i), S(commonChars(i)), S(o))

  i, o = []string{"cool","lock","cook"}, []string{"c","o"}
  T(t, S(i), S(commonChars(i)), S(o))
}

