package main
import "testing"
import . "main/pkg/testing_helper"

func TestAddBinary(_t *testing.T) {
  a, b, o := "11", "1", "100"
  T(_t, S(a, b), S(addBinary(a, b)), S(o))

  a, b, o = "1010", "1011", "10101"
  T(_t, S(a, b), S(addBinary(a, b)), S(o))
}
