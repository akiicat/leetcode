package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindMaxConsecutiveOnes(_t *testing.T) {
  i, o := []int{1,1,0,1,1,1}, 3
  T(_t, S(i), S(findMaxConsecutiveOnes(i)), S(o))

  i, o = []int{1,0,1,1,0,1}, 2
  T(_t, S(i), S(findMaxConsecutiveOnes(i)), S(o))
}

func TestFindMaxConsecutiveOnesCounter(_t *testing.T) {
  i, o := []int{1,1,0,1,1,1}, 3
  T(_t, S(i), S(findMaxConsecutiveOnesCounter(i)), S(o))

  i, o = []int{1,0,1,1,0,1}, 2
  T(_t, S(i), S(findMaxConsecutiveOnesCounter(i)), S(o))
}


