package main
import "testing"
import . "main/pkg/testing_helper"

func TestSortString(_t *testing.T) {
  i, o := "aaaabbbbcccc", "abccbaabccba"
  T(_t, S(i), S(sortString(i)), S(o))

  i, o = "rat", "art"
  T(_t, S(i), S(sortString(i)), S(o))

  i, o = "leetcode", "cdelotee"
  T(_t, S(i), S(sortString(i)), S(o))

  i, o = "ggggggg", "ggggggg"
  T(_t, S(i), S(sortString(i)), S(o))

  i, o = "spo", "ops"
  T(_t, S(i), S(sortString(i)), S(o))
}

