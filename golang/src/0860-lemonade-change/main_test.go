package main
import "testing"
import . "main/pkg/testing_helper"

func TestLemonadeChange(_t *testing.T) {
  i, o := []int{5,5,5,10,20}, true
  T(_t, S(i), S(lemonadeChange(i)), S(o))

  i, o = []int{5,5,10}, true
  T(_t, S(i), S(lemonadeChange(i)), S(o))

  i, o = []int{10,10}, false
  T(_t, S(i), S(lemonadeChange(i)), S(o))

  i, o = []int{5,5,10,10,20}, false
  T(_t, S(i), S(lemonadeChange(i)), S(o))
}

