package main
import "testing"
import . "main/pkg/testing_helper"

func TestCompress(_t *testing.T) {
  c, o1, o2 := []byte{'a','a','b','b','c','c','c'}, 6, "a2b2c3"
  T(_t, S(c), S(compress(c), c[:o1]), S(o1, o2))

  c, o1, o2 = []byte{'a'}, 1, "a"
  T(_t, S(c), S(compress(c), c[:o1]), S(o1, o2))

  c, o1, o2 = []byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'}, 4, "ab12"
  T(_t, S(c), S(compress(c), c[:o1]), S(o1, o2))
}

