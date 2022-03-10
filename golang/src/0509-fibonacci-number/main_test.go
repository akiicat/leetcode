package main
import "testing"
import . "main/pkg/testing_helper"

func TestFib(_t *testing.T) {
  i, o := 2, 1
  T(_t, S(i), S(fib(i)), S(o))

  i, o = 3, 2
  T(_t, S(i), S(fib(i)), S(o))

  i, o = 4, 3
  T(_t, S(i), S(fib(i)), S(o))

  i, o = 5, 5
  T(_t, S(i), S(fib(i)), S(o))

  i, o = 9, 34
  T(_t, S(i), S(fib(i)), S(o))

  i, o = 10, 55
  T(_t, S(i), S(fib(i)), S(o))
}

func TestFibFormula(_t *testing.T) {
  i, o := 2, 1
  T(_t, S(i), S(fibFormula(i)), S(o))

  i, o = 3, 2
  T(_t, S(i), S(fibFormula(i)), S(o))

  i, o = 4, 3
  T(_t, S(i), S(fibFormula(i)), S(o))

  i, o = 5, 5
  T(_t, S(i), S(fibFormula(i)), S(o))

  i, o = 9, 34
  T(_t, S(i), S(fibFormula(i)), S(o))

  i, o = 10, 55
  T(_t, S(i), S(fibFormula(i)), S(o))
}

func TestFibRecursive(_t *testing.T) {
  i, o := 2, 1
  T(_t, S(i), S(fibRecursive(i)), S(o))

  i, o = 3, 2
  T(_t, S(i), S(fibRecursive(i)), S(o))

  i, o = 4, 3
  T(_t, S(i), S(fibRecursive(i)), S(o))

  i, o = 5, 5
  T(_t, S(i), S(fibRecursive(i)), S(o))

  i, o = 9, 34
  T(_t, S(i), S(fibRecursive(i)), S(o))

  i, o = 10, 55
  T(_t, S(i), S(fibRecursive(i)), S(o))
}

