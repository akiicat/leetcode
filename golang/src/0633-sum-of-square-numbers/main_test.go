package main
import "testing"
import . "main/pkg/testing_helper"

func TestJudgeSquareSum(_t *testing.T) {
  i, o := 2, true
  T(_t, S(i), S(judgeSquareSum(i)), S(o))

  i, o = 3, false
  T(_t, S(i), S(judgeSquareSum(i)), S(o))

  i, o = 5, true
  T(_t, S(i), S(judgeSquareSum(i)), S(o))

  i, o = 13, true
  T(_t, S(i), S(judgeSquareSum(i)), S(o))

  i, o = 2147483645, false
  T(_t, S(i), S(judgeSquareSum(i)), S(o))
}

func TestJudgeSquareSumSqrt(_t *testing.T) {
  i, o := 2, true
  T(_t, S(i), S(judgeSquareSumSqrt(i)), S(o))

  i, o = 3, false
  T(_t, S(i), S(judgeSquareSumSqrt(i)), S(o))

  i, o = 5, true
  T(_t, S(i), S(judgeSquareSumSqrt(i)), S(o))

  i, o = 13, true
  T(_t, S(i), S(judgeSquareSumSqrt(i)), S(o))

  i, o = 2147483645, false
  T(_t, S(i), S(judgeSquareSumSqrt(i)), S(o))
}

func TestJudgeSquareSumBetterBruteForce(_t *testing.T) {
  i, o := 2, true
  T(_t, S(i), S(judgeSquareSumBetterBruteForce(i)), S(o))

  i, o = 3, false
  T(_t, S(i), S(judgeSquareSumBetterBruteForce(i)), S(o))

  i, o = 5, true
  T(_t, S(i), S(judgeSquareSumBetterBruteForce(i)), S(o))

  i, o = 13, true
  T(_t, S(i), S(judgeSquareSumBetterBruteForce(i)), S(o))

  i, o = 2147483645, false
  T(_t, S(i), S(judgeSquareSumBetterBruteForce(i)), S(o))
}

func TestJudgeSquareSumBruteForce(_t *testing.T) {
  i, o := 2, true
  T(_t, S(i), S(judgeSquareSumBruteForce(i)), S(o))

  i, o = 3, false
  T(_t, S(i), S(judgeSquareSumBruteForce(i)), S(o))

  i, o = 5, true
  T(_t, S(i), S(judgeSquareSumBruteForce(i)), S(o))

  i, o = 13, true
  T(_t, S(i), S(judgeSquareSumBruteForce(i)), S(o))

  i, o = 2147483645, false
  T(_t, S(i), S(judgeSquareSumBruteForce(i)), S(o))
}

func TestJudgeSquareSumBinarySearch(_t *testing.T) {
  i, o := 2, true
  T(_t, S(i), S(judgeSquareSumBinarySearch(i)), S(o))

  i, o = 3, false
  T(_t, S(i), S(judgeSquareSumBinarySearch(i)), S(o))

  i, o = 5, true
  T(_t, S(i), S(judgeSquareSumBinarySearch(i)), S(o))

  i, o = 13, true
  T(_t, S(i), S(judgeSquareSumBinarySearch(i)), S(o))

  i, o = 2147483645, false
  T(_t, S(i), S(judgeSquareSumBinarySearch(i)), S(o))
}

func TestJudgeSquareSumFermatTheorem(_t *testing.T) {
  i, o := 2, true
  T(_t, S(i), S(judgeSquareSumFermatTheorem(i)), S(o))

  i, o = 3, false
  T(_t, S(i), S(judgeSquareSumFermatTheorem(i)), S(o))

  i, o = 5, true
  T(_t, S(i), S(judgeSquareSumFermatTheorem(i)), S(o))

  i, o = 13, true
  T(_t, S(i), S(judgeSquareSumFermatTheorem(i)), S(o))

  i, o = 2147483645, false
  T(_t, S(i), S(judgeSquareSumFermatTheorem(i)), S(o))
}

