package main
import "testing"
import . "main/pkg/testing_helper"

func TestCountNegatives(_t *testing.T) {
  i, o := [][]int{
    []int{4,3,2,-1},
    []int{3,2,1,-1},
    []int{1,1,-1,-2},
    []int{-1,-1,-2,-3},
  }, 8
  T(_t, S(i), S(countNegatives(i)), S(o))

  i, o = [][]int{
    []int{3,2},
    []int{1,0},
  }, 0
  T(_t, S(i), S(countNegatives(i)), S(o))

  i, o = [][]int{
    []int{1,-1},
    []int{-1,-1},
  }, 3
  T(_t, S(i), S(countNegatives(i)), S(o))

  i, o = [][]int{
    []int{1},
  }, 0
  T(_t, S(i), S(countNegatives(i)), S(o))
}

func TestCountNegativesRow(_t *testing.T) {
  i, o := [][]int{
    []int{4,3,2,-1},
    []int{3,2,1,-1},
    []int{1,1,-1,-2},
    []int{-1,-1,-2,-3},
  }, 8
  T(_t, S(i), S(countNegativesRow(i)), S(o))

  i, o = [][]int{
    []int{3,2},
    []int{1,0},
  }, 0
  T(_t, S(i), S(countNegativesRow(i)), S(o))

  i, o = [][]int{
    []int{1,-1},
    []int{-1,-1},
  }, 3
  T(_t, S(i), S(countNegativesRow(i)), S(o))

  i, o = [][]int{
    []int{1},
  }, 0
  T(_t, S(i), S(countNegativesRow(i)), S(o))
}

