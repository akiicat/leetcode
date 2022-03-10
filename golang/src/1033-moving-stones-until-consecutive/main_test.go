package main
import "testing"
import . "main/pkg/testing_helper"

func TestNumMovesStones(_t *testing.T) {
  a, b, c, o := 1, 2, 5, []int{1,2}
  T(_t, S(a, b, c), S(numMovesStones(a, b, c)), S(o))

  a, b, c, o = 4, 3, 2, []int{0,0}
  T(_t, S(a, b, c), S(numMovesStones(a, b, c)), S(o))

  a, b, c, o = 3, 5, 1, []int{1,2}
  T(_t, S(a, b, c), S(numMovesStones(a, b, c)), S(o))

  a, b, c, o = 7, 4, 1, []int{2,4}
  T(_t, S(a, b, c), S(numMovesStones(a, b, c)), S(o))
}

