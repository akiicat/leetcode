package main
import "testing"
import . "main/pkg/testing_helper"

func TestMinSubsequence(_t *testing.T) {
  i, o := []int{4,3,10,9,8}, []int{10,9}
  T(_t, S(i), S(minSubsequence(i)), S(o))

  i, o = []int{4,4,7,6,7}, []int{7,7,6}
  T(_t, S(i), S(minSubsequence(i)), S(o))

  i, o = []int{6}, []int{6}
  T(_t, S(i), S(minSubsequence(i)), S(o))
}

