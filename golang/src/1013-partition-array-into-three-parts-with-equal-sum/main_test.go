package main
import "testing"
import . "main/pkg/testing_helper"

func TestCanThreePartsEqualSum(t *testing.T) {
  i, o := []int{0,2,1,-6,6,-7,9,1,2,0,1}, true
  T(t, S(i), S(canThreePartsEqualSum(i)), S(o))

  i, o = []int{0,2,1,-6,6,7,9,-1,2,0,1}, false
  T(t, S(i), S(canThreePartsEqualSum(i)), S(o))

  i, o = []int{3,3,6,5,-2,2,5,1,-9,4}, true
  T(t, S(i), S(canThreePartsEqualSum(i)), S(o))
}

