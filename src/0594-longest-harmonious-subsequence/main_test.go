package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindLHS(_t *testing.T) {
  i, o := []int{1,3,2,2,5,2,3,7}, 5
  T(_t, S(i), S(findLHS(i)), S(o))

  i, o = []int{1,1,1,1}, 0
  T(_t, S(i), S(findLHS(i)), S(o))
}
