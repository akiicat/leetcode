package main
import "testing"
import . "main/pkg/testing_helper"

func TestContainsNearbyDuplicate(_t *testing.T) {
  i, k, o := []int{1,2,3,1}, 3, true
  T(_t, S(i, k), S(containsNearbyDuplicate(i, k)), S(o))

  i, k, o = []int{1,0,1,1}, 1, true
  T(_t, S(i, k), S(containsNearbyDuplicate(i, k)), S(o))

  i, k, o = []int{1,2,3,1,2,3}, 2, false
  T(_t, S(i, k), S(containsNearbyDuplicate(i, k)), S(o))
}
