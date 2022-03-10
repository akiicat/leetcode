package main
import "testing"
import . "main/pkg/testing_helper"

func TestSmallerNumbersThanCurrent(_t *testing.T) {
  i, o := []int{8,1,2,2,3}, []int{4,0,1,1,3}
  T(_t, S(i), S(smallerNumbersThanCurrent(i)), S(o))

  i, o = []int{6,5,4,8}, []int{2,1,0,3}
  T(_t, S(i), S(smallerNumbersThanCurrent(i)), S(o))

  i, o = []int{7,7,7,7}, []int{0,0,0,0}
  T(_t, S(i), S(smallerNumbersThanCurrent(i)), S(o))
}

