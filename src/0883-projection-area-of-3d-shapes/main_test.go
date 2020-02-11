package main
import "testing"
import . "main/pkg/testing_helper"

func TestProjectionArea(t *testing.T) {
  i, o := [][]int{
    []int{2},
  }, 5
  T(t, S(i), S(projectionArea(i)), S(o))

  i, o = [][]int{
    []int{1,2},
    []int{3,4},
  }, 17
  T(t, S(i), S(projectionArea(i)), S(o))

  i, o = [][]int{
    []int{1,0},
    []int{0,2},
  }, 8
  T(t, S(i), S(projectionArea(i)), S(o))

  i, o = [][]int{
    []int{1,1,1},
    []int{1,0,1},
    []int{1,1,1},
  }, 14
  T(t, S(i), S(projectionArea(i)), S(o))

  i, o = [][]int{
    []int{2,2,2},
    []int{2,1,2},
    []int{2,2,2},
  }, 21
  T(t, S(i), S(projectionArea(i)), S(o))
}

