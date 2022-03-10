package main
import "testing"
import . "main/pkg/testing_helper"

func TestMaxNumberOfBalloons(t *testing.T) {
  i, o := "nlaebolko", 1
  T(t, S(i), S(maxNumberOfBalloons(i)), S(o))

  i, o = "loonbalxballpoon", 2
  T(t, S(i), S(maxNumberOfBalloons(i)), S(o))

  i, o = "leetcode", 0
  T(t, S(i), S(maxNumberOfBalloons(i)), S(o))
}

