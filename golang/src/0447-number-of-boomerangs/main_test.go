package main
import "testing"
import . "main/pkg/testing_helper"

func TestNumberOfBoomerangs(_t *testing.T) {
  i, o := [][]int{[]int{0,0},[]int{1,0},[]int{2,0}}, 2
  T(_t, S(i), S(numberOfBoomerangs(i)), S(o))

  i, o = [][]int{[]int{0,0},[]int{1,0},[]int{-1,0},[]int{0,1},[]int{0,-1}}, 20
  T(_t, S(i), S(numberOfBoomerangs(i)), S(o))
}

func TestNumberOfBoomerangsNested(_t *testing.T) {
  i, o := [][]int{[]int{0,0},[]int{1,0},[]int{2,0}}, 2
  T(_t, S(i), S(numberOfBoomerangsNested(i)), S(o))

  i, o = [][]int{[]int{0,0},[]int{1,0},[]int{-1,0},[]int{0,1},[]int{0,-1}}, 20
  T(_t, S(i), S(numberOfBoomerangsNested(i)), S(o))
}
