package main
import "testing"
import . "main/pkg/testing_helper"

func TestGenerateParenthesis(_t *testing.T) {
  i, o := 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
  T(_t, S(i), S(generateParenthesis(i)), S(o))
}

