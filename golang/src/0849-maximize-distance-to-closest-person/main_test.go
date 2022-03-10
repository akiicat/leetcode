package main
import "testing"
import . "main/pkg/testing_helper"

func TestMaxDistToClosest(_t *testing.T) {
  i, o := []int{1,0,0,0,1,0,1}, 2
  T(_t, S(i), S(maxDistToClosest(i)), S(o))

  i, o = []int{1,0,0,0,0,1,0,1}, 2
  T(_t, S(i), S(maxDistToClosest(i)), S(o))

  i, o = []int{1,0,0,0}, 3
  T(_t, S(i), S(maxDistToClosest(i)), S(o))

  i, o = []int{0,0,1,0}, 2
  T(_t, S(i), S(maxDistToClosest(i)), S(o))
}

