package main
import "testing"
import . "main/pkg/testing_helper"

func TestMinCostClimbingStairs(_t *testing.T) {
  i, o := []int{10,15,20}, 15
  T(_t, S(i), S(minCostClimbingStairs(i)), S(o))

  i, o = []int{1,100,1,1,1,100,1,1,100,1}, 6
  T(_t, S(i), S(minCostClimbingStairs(i)), S(o))
}

func TestMinCostClimbingStairsRecursive(_t *testing.T) {
  i, o := []int{10,15,20}, 15
  T(_t, S(i), S(minCostClimbingStairsRecursive(i)), S(o))

  i, o = []int{1,100,1,1,1,100,1,1,100,1}, 6
  T(_t, S(i), S(minCostClimbingStairsRecursive(i)), S(o))
}

