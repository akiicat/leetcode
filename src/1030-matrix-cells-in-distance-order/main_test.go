package main
import "testing"
import . "main/pkg/testing_helper"

func TestAllCellsDistOrder(_t *testing.T) {
  r, c, r0, c0, o := 1, 2, 0, 0, [][]int{[]int{0,0},[]int{0,1}}
  T(_t, S(r, c, r0, c0), S(allCellsDistOrder(r, c, r0, c0)), S(o))

  r, c, r0, c0, o = 2, 2, 0, 1, [][]int{[]int{0,1},[]int{1,1},[]int{0,0},[]int{1,0}}
  T(_t, S(r, c, r0, c0), S(allCellsDistOrder(r, c, r0, c0)), S(o))

  r, c, r0, c0, o = 2, 3, 1, 2, [][]int{[]int{1,2},[]int{0,2},[]int{1,1},[]int{0,1},[]int{1,0},[]int{0,0}}
  T(_t, S(r, c, r0, c0), S(allCellsDistOrder(r, c, r0, c0)), S(o))
}

