package main
import "testing"
import . "main/pkg/testing_helper"

func TestReverseString(_t *testing.T) {
  i, o := []byte{'h','e','l','l','o'}, []byte{'o','l','l','e','h'}
  _i := S(i)
  reverseString(i)
  T(_t, _i, S(i), S(o))
}

func TestReverseStringTwoIndex(_t *testing.T) {
  i, o := []byte{'h','e','l','l','o'}, []byte{'o','l','l','e','h'}
  _i := S(i)
  reverseStringTwoIndex(i)
  T(_t, _i, S(i), S(o))
}
