package main
import "testing"
import . "main/pkg/testing_helper"

func TestDominantIndex(_t *testing.T) {
  i, o := []int{3,6,1,0}, 1
  T(_t, S(i), S(dominantIndex(i)), S(o))

  i, o = []int{1,2,3,4}, -1
  T(_t, S(i), S(dominantIndex(i)), S(o))
}

