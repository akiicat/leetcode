package main
import "testing"
import . "main/pkg/testing_helper"

func TestMostCommonWord(_t *testing.T) {
  p, b, o := "Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}, "ball"
  T(_t, S(p, b), S(mostCommonWord(p, b)), S(o))

  p, b, o = "Bob", []string{}, "bob"
  T(_t, S(p, b), S(mostCommonWord(p, b)), S(o))
}

