package main
import "testing"
import . "main/pkg/testing_helper"

func TestCountAndSay(_t *testing.T) {
  i, o := 1, "1"
  T(_t, S(i), S(countAndSay(i)), S(o))

  i, o = 2, "11"
  T(_t, S(i), S(countAndSay(i)), S(o))

  i, o = 3, "21"
  T(_t, S(i), S(countAndSay(i)), S(o))

  i, o = 4, "1211"
  T(_t, S(i), S(countAndSay(i)), S(o))

  i, o = 5, "111221"
  T(_t, S(i), S(countAndSay(i)), S(o))
}

func TestCountAndSayRecursive(_t *testing.T) {
  i, o := 1, "1"
  T(_t, S(i), S(countAndSayRecursive(i)), S(o))

  i, o = 2, "11"
  T(_t, S(i), S(countAndSayRecursive(i)), S(o))

  i, o = 3, "21"
  T(_t, S(i), S(countAndSayRecursive(i)), S(o))

  i, o = 4, "1211"
  T(_t, S(i), S(countAndSayRecursive(i)), S(o))

  i, o = 5, "111221"
  T(_t, S(i), S(countAndSayRecursive(i)), S(o))
}

