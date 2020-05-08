package main
import "testing"
import . "main/pkg/testing_helper"

func TestMaskPII(_t *testing.T) {
  i, o := "LeetCode@LeetCode.com", "l*****e@leetcode.com"
  T(_t, S(i), S(maskPII(i)), S(o))

  i, o = "AB@qq.com", "a*****b@qq.com"
  T(_t, S(i), S(maskPII(i)), S(o))

  i, o = "1(234)567-890", "***-***-7890"
  T(_t, S(i), S(maskPII(i)), S(o))

  i, o = "86-(10)12345678", "+**-***-***-5678"
  T(_t, S(i), S(maskPII(i)), S(o))
}

