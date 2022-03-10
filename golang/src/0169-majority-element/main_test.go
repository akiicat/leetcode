package main
import "testing"
import . "main/pkg/testing_helper"

func TestMajorityElement(_t *testing.T) {
  i, o := []int{3,2,3}, 3
  T(_t, S(i), S(majorityElement(i)), S(o))

  i, o = []int{2,2,1,1,1,2,2}, 2
  T(_t, S(i), S(majorityElement(i)), S(o))
}

func TestMajorityElementMap(_t *testing.T) {
  i, o := []int{3,2,3}, 3
  T(_t, S(i), S(majorityElementMap(i)), S(o))

  i, o = []int{2,2,1,1,1,2,2}, 2
  T(_t, S(i), S(majorityElementMap(i)), S(o))
}
