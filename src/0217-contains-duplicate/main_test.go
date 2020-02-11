package main
import "testing"
import . "main/pkg/testing_helper"

func TestContainsDuplicate(_t *testing.T) {
  i, o := []int{1,2,3,1}, true
  T(_t, S(i), S(containsDuplicate(i)), S(o))

  i, o = []int{1,2,3,4}, false
  T(_t, S(i), S(containsDuplicate(i)), S(o))

  i, o = []int{1,1,1,3,3,4,3,2,4,2}, true
  T(_t, S(i), S(containsDuplicate(i)), S(o))
}
