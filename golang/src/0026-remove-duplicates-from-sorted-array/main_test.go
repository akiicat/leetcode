package main
import "testing"
import . "main/pkg/testing_helper"

func TestRemoveDuplicates(_t *testing.T) {
  i, o1, o2 := []int{1,1,2}, []int{1,2}, 2
  T(_t, S(i), S(i[0:o2], removeDuplicates(i)), S(o1, o2))

  i, o1, o2 = []int{0,0,1,1,1,2,2,3,3,4}, []int{0,1,2,3,4}, 5
  T(_t, S(i), S(i[0:o2], removeDuplicates(i)), S(o1, o2))
}
