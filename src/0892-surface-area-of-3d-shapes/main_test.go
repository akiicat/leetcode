package main
import "testing"
import . "main/pkg/testing_helper"

func TestSurfaceArea(t *testing.T) {
  i, o := [][]int{
    []int{2},
  }, 10
  T(t, S(i), S(surfaceArea(i)), S(o))

  i, o = [][]int{
    []int{1,2},
    []int{3,4},
  }, 34
  T(t, S(i), S(surfaceArea(i)), S(o))

  i, o = [][]int{
    []int{1,0},
    []int{0,2},
  }, 16
  T(t, S(i), S(surfaceArea(i)), S(o))

  i, o = [][]int{
    []int{1,1,1},
    []int{1,0,1},
    []int{1,1,1},
  }, 32
  T(t, S(i), S(surfaceArea(i)), S(o))

  i, o = [][]int{
    []int{2,2,2},
    []int{2,1,2},
    []int{2,2,2},
  }, 46
  T(t, S(i), S(surfaceArea(i)), S(o))
}

