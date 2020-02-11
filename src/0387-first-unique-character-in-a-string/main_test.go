package main
import "testing"
import . "main/pkg/testing_helper"

func TestFirstUniqChar(_t *testing.T) {
  i, o := "leetcode", 0
  T(_t, S(i), S(firstUniqChar(i)), S(o))

  i, o = "loveleetcode", 2
  T(_t, S(i), S(firstUniqChar(i)), S(o))

  i, o = "lel", 1
  T(_t, S(i), S(firstUniqChar(i)), S(o))

  i, o = "cc", -1
  T(_t, S(i), S(firstUniqChar(i)), S(o))
}

func TestFirstUniqCharDiscreteTransform(_t *testing.T) {
  i, o := "leetcode", 0
  T(_t, S(i), S(firstUniqCharDiscreteTransform(i)), S(o))

  i, o = "loveleetcode", 2
  T(_t, S(i), S(firstUniqCharDiscreteTransform(i)), S(o))

  i, o = "lel", 1
  T(_t, S(i), S(firstUniqCharDiscreteTransform(i)), S(o))

  i, o = "cc", -1
  T(_t, S(i), S(firstUniqCharDiscreteTransform(i)), S(o))
}
