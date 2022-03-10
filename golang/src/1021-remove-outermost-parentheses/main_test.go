package main
import "testing"
import . "main/pkg/testing_helper"

func TestRemoveOuterParentheses(t *testing.T) {
  i, o := "(()())(())", "()()()"
  T(t, S(i), S(removeOuterParentheses(i)), S(o))

  i, o = "(()())(())(()(()))", "()()()()(())"
  T(t, S(i), S(removeOuterParentheses(i)), S(o))

  i, o = "()()", ""
  T(t, S(i), S(removeOuterParentheses(i)), S(o))
}

