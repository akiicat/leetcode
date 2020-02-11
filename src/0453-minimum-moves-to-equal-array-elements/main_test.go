package main
import "testing"
import . "main/pkg/testing_helper"

func TestMinMoves(_t *testing.T) {
  i, o := []int{1,2,3}, 3
  T(_t, S(i), S(minMoves(i)), S(o))

  i, o = []int{1,2,8}, 8
  T(_t, S(i), S(minMoves(i)), S(o))
}

func TestMinMovesDiff(_t *testing.T) {
  i, o := []int{1,2,3}, 3
  T(_t, S(i), S(minMovesDiff(i)), S(o))

  i, o = []int{1,2,8}, 8
  T(_t, S(i), S(minMovesDiff(i)), S(o))
}
