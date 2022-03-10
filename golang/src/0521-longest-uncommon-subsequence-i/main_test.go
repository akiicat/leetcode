package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindLUSlength(_t *testing.T) {
  i1, i2, o := "abc", "abc", -1
  T(_t, S(i1, i2), S(findLUSlength(i1, i2)), S(o))

  i1, i2, o = "abc", "cbc", 3
  T(_t, S(i1, i2), S(findLUSlength(i1, i2)), S(o))
}

