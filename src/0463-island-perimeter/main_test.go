package main
import "testing"
import . "main/pkg/testing_helper"

func TestIslandPerimeter(_t *testing.T) {
  i, o := [][]int{[]int{0,1,0,0}, []int{1,1,1,0}, []int{0,1,0,0}, []int{1,1,0,0}}, 16
  T(_t, S(i), S(islandPerimeter(i)), S(o))

  i, o = [][]int{[]int{1,0}}, 4
  T(_t, S(i), S(islandPerimeter(i)), S(o))
}

func TestIslandPerimeterRemove(_t *testing.T) {
  i, o := [][]int{[]int{0,1,0,0}, []int{1,1,1,0}, []int{0,1,0,0}, []int{1,1,0,0}}, 16
  T(_t, S(i), S(islandPerimeterRemove(i)), S(o))

  i, o = [][]int{[]int{1,0}}, 4
  T(_t, S(i), S(islandPerimeterRemove(i)), S(o))
}
