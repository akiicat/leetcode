package main
import "testing"
import . "main/pkg/testing_helper"

func TestRemoveDuplicates(_t *testing.T) {
  i, n, o := []int{3,2,2,3}, 3, 2
  T(_t, S(i, n), S(removeElement(i, n)), S(o))

  i, n, o = []int{0,1,2,2,3,0,4,2}, 2, 5
  T(_t, S(i, n), S(removeElement(i, n)), S(o))
}

