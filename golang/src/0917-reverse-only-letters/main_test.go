package main
import "testing"
import . "main/pkg/testing_helper"

func TestReverseOnlyLetters(_t *testing.T) {
  i, o := "ab-cd", "dc-ba"
  T(_t, S(i), S(reverseOnlyLetters(i)), S(o))

  i, o = "a-bC-dEf-ghIj", "j-Ih-gfE-dCba"
  T(_t, S(i), S(reverseOnlyLetters(i)), S(o))

  i, o = "Test1ng-Leet=code-Q!", "Qedo1ct-eeLg=ntse-T!"
  T(_t, S(i), S(reverseOnlyLetters(i)), S(o))

  i, o = "7_28]", "7_28]"
  T(_t, S(i), S(reverseOnlyLetters(i)), S(o))
}
