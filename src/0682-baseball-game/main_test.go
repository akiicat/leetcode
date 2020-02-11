package main
import "testing"
import . "main/pkg/testing_helper"

func TestCalPoints(_t *testing.T) {
  i, o := []string{"5","2","C","D","+"}, 30
  T(_t, S(i), S(calPoints(i)), S(o))

  i, o = []string{"5","-2","4","C","D","9","+","+"}, 27
  T(_t, S(i), S(calPoints(i)), S(o))
}

