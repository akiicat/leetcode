package main
import "testing"
import . "main/pkg/testing_helper"

func TestLetterCombinations(_t *testing.T) {
  i, o := "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
  T(_t, S(i), S(Sort(letterCombinations(i))), S(Sort(o)))

  i, o = "", []string{}
  T(_t, S(i), S(Sort(letterCombinations(i))), S(Sort(o)))
}

