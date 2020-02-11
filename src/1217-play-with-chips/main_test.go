package main
import "testing"
import . "main/pkg/testing_helper"

func TestMinCostToMoveChips(t *testing.T) {
  i, o := []int{1,2,3}, 1
  T(t, S(i), S(minCostToMoveChips(i)), S(o))

  i, o = []int{2,2,2,3,3}, 2
  T(t, S(i), S(minCostToMoveChips(i)), S(o))

  i, o = []int{1,2,2,2,2}, 1
  T(t, S(i), S(minCostToMoveChips(i)), S(o))
}

