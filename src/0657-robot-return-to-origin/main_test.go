package main
import "testing"
import . "main/pkg/testing_helper"

func TestJudgeCircle(_t *testing.T) {
  i, o := "UD", true
  T(_t, S(i), S(judgeCircle(i)), S(o))

  i, o = "LL", false
  T(_t, S(i), S(judgeCircle(i)), S(o))
}
