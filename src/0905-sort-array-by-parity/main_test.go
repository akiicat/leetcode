package main
import "testing"
import . "main/pkg/testing_helper"

func TestSortArrayByParity(t *testing.T) {
  i, o := []int{3,1,2,4}, []int{2,4,3,1}
  T(t, S(i), S(sortArrayByParity(i)), S(o))
}

