package main
import "testing"
import . "main/pkg/testing_helper"

func TestOrangesRotting(_t *testing.T) {
  i, o := [][]int{[]int{2,1,1},[]int{1,1,0},[]int{0,1,1}}, 4
  T(_t, S(i), S(orangesRotting(i)), S(o))

  i, o = [][]int{[]int{2,1,1},[]int{0,1,1},[]int{1,0,1}}, -1
  T(_t, S(i), S(orangesRotting(i)), S(o))

  i, o = [][]int{[]int{0,2}}, 0
  T(_t, S(i), S(orangesRotting(i)), S(o))

  i, o = [][]int{[]int{0}}, 0
  T(_t, S(i), S(orangesRotting(i)), S(o))
}


func TestOrangesRottingBFS(_t *testing.T) {
  i, o := [][]int{[]int{2,1,1},[]int{1,1,0},[]int{0,1,1}}, 4
  T(_t, S(i), S(orangesRottingBFS(i)), S(o))

  i, o = [][]int{[]int{2,1,1},[]int{0,1,1},[]int{1,0,1}}, -1
  T(_t, S(i), S(orangesRottingBFS(i)), S(o))

  i, o = [][]int{[]int{0,2}}, 0
  T(_t, S(i), S(orangesRottingBFS(i)), S(o))

  i, o = [][]int{[]int{0}}, 0
  T(_t, S(i), S(orangesRottingBFS(i)), S(o))
}
