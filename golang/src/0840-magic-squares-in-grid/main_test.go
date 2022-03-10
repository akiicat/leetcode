package main
import "testing"
import . "main/pkg/testing_helper"

func TestNumMagicSquaresInside(_t *testing.T) {
  i, o := [][]int{
    []int{4,3,8,4},
    []int{9,5,1,9},
    []int{2,7,6,2},
  }, 1
  T(_t, S(i), S(numMagicSquaresInside(i)), S(o))

  i, o = [][]int{
    []int{5,5,5},
    []int{5,5,5},
    []int{5,5,5},
  }, 0
  T(_t, S(i), S(numMagicSquaresInside(i)), S(o))

  i, o = [][]int{
    []int{3,2,1,6},
    []int{5,9,6,8},
    []int{1,5,1,2},
    []int{3,7,3,4},
  }, 0
  T(_t, S(i), S(numMagicSquaresInside(i)), S(o))

  i, o = [][]int{
    []int{1,3,7,8,8},
    []int{8,3,2,7,4},
    []int{3,8,4,0,9},
    []int{8,1,6,5,0},
    []int{7,2,1,8,6},
  }, 0
  T(_t, S(i), S(numMagicSquaresInside(i)), S(o))
}

