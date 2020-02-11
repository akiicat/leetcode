package main
import "testing"
import . "main/pkg/testing_helper"

func TestDistributeCandies(_t *testing.T) {
  i, o := []int{1,1,2,2,3,3}, 3
  T(_t, S(i), S(distributeCandies(i)), S(o))

  i, o = []int{1,1,2,3}, 2
  T(_t, S(i), S(distributeCandies(i)), S(o))
}

