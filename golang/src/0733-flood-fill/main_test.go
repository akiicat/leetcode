package main
import "testing"
import . "main/pkg/testing_helper"

func TestFloodFill(_t *testing.T) {
  image, sr, sc, newColor, o := [][]int{[]int{1,1,1}, []int{1,1,0}, []int{1,0,1}}, 1, 1, 2, [][]int{[]int{2,2,2}, []int{2,2,0}, []int{2,0,1}}
  T(_t, S(image, sr, sc, newColor), S(floodFill(image, sr, sc, newColor)), S(o))

  image, sr, sc, newColor, o = [][]int{[]int{0,0,0}, []int{0,1,1}}, 1, 1, 1, [][]int{[]int{0,0,0}, []int{0,1,1}}
  T(_t, S(image, sr, sc, newColor), S(floodFill(image, sr, sc, newColor)), S(o))
}

