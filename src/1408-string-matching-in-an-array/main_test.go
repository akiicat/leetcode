package main
import "testing"
import . "main/pkg/testing_helper"

func TestStringMatching(_t *testing.T) {
  i, o := []string{"mass","as","hero","superhero"}, []string{"as","hero"}
  T(_t, S(i), S(Sort(stringMatching(i))), S(Sort(o)))

  i, o = []string{"leetcode","et","code"}, []string{"et","code"}
  T(_t, S(i), S(Sort(stringMatching(i))), S(Sort(o)))

  i, o = []string{"blue","green","bu"}, []string{}
  T(_t, S(i), S(Sort(stringMatching(i))), S(Sort(o)))
}

