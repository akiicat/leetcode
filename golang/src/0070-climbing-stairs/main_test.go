package main
import "testing"
import . "main/pkg/testing_helper"

func TestClimbStairs(_t *testing.T) {
  i, o := 2, 2
  T(_t, S(i), S(climbStairs(i)), S(o))

  i, o = 3, 3
  T(_t, S(i), S(climbStairs(i)), S(o))

  i, o = 4, 5
  T(_t, S(i), S(climbStairs(i)), S(o))

  i, o = 5, 8
  T(_t, S(i), S(climbStairs(i)), S(o))

  i, o = 9, 55
  T(_t, S(i), S(climbStairs(i)), S(o))
}

func TestClimbStairsFormula(_t *testing.T) {
  i, o := 2, 2
  T(_t, S(i), S(climbStairsFormula(i)), S(o))

  i, o = 3, 3
  T(_t, S(i), S(climbStairsFormula(i)), S(o))

  i, o = 4, 5
  T(_t, S(i), S(climbStairsFormula(i)), S(o))

  i, o = 5, 8
  T(_t, S(i), S(climbStairsFormula(i)), S(o))

  i, o = 9, 55
  T(_t, S(i), S(climbStairsFormula(i)), S(o))
}

func TestClimbStairsRecursive(_t *testing.T) {
  i, o := 2, 2
  T(_t, S(i), S(climbStairsRecursive(i)), S(o))

  i, o = 3, 3
  T(_t, S(i), S(climbStairsRecursive(i)), S(o))

  i, o = 4, 5
  T(_t, S(i), S(climbStairsRecursive(i)), S(o))

  i, o = 5, 8
  T(_t, S(i), S(climbStairsRecursive(i)), S(o))

  i, o = 9, 55
  T(_t, S(i), S(climbStairsRecursive(i)), S(o))
}
