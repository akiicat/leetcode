package main
import "testing"
import . "main/pkg/testing_helper"

func TestOddCells(_t *testing.T) {
  n, m, i, o := 2, 3, [][]int{[]int{0,1},[]int{1,1}}, 6
  T(_t, S(n, m, i), S(oddCells(n, m, i)), S(o))

  n, m, i, o = 2, 2, [][]int{[]int{1,1},[]int{0,0}}, 0
  T(_t, S(n, m, i), S(oddCells(n, m, i)), S(o))
}

