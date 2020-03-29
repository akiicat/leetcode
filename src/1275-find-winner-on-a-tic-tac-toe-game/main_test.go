package main
import "testing"
import . "main/pkg/testing_helper"

func TestTictactoe(_t *testing.T) {
  i, o := [][]int{[]int{0,0},[]int{2,0},[]int{1,1},[]int{2,1},[]int{2,2}}, "A"
  T(_t, S(i), S(tictactoe(i)), S(o))

  i, o = [][]int{[]int{0,0},[]int{1,1},[]int{0,1},[]int{0,2},[]int{1,0},[]int{2,0}}, "B"
  T(_t, S(i), S(tictactoe(i)), S(o))

  i, o = [][]int{[]int{0,0},[]int{1,1},[]int{2,0},[]int{1,0},[]int{1,2},[]int{2,1},[]int{0,1},[]int{0,2},[]int{2,2}}, "Draw"
  T(_t, S(i), S(tictactoe(i)), S(o))

  i, o = [][]int{[]int{0,0},[]int{1,1}}, "Pending"
  T(_t, S(i), S(tictactoe(i)), S(o))
}

