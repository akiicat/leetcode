package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindRelativeRanks(_t *testing.T) {
  i, o := []int{5,4,3,2,1}, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}
  T(_t, S(i), S(findRelativeRanks(i)), S(o))
}

