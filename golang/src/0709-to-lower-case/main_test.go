package main
import "testing"
import . "main/pkg/testing_helper"

func TestToLowerCase(_t *testing.T) {
  i, o := "Hello", "hello"
  T(_t, S(i), S(toLowerCase(i)), S(o))

  i, o = "here", "here"
  T(_t, S(i), S(toLowerCase(i)), S(o))

  i, o = "LOVELY", "lovely"
  T(_t, S(i), S(toLowerCase(i)), S(o))
}
