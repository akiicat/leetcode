package main
import "testing"
import . "main/pkg/testing_helper"

func TestMatrixReshape(_t *testing.T) {
  i, r, c, o := [][]int{[]int{1,2}, []int{3,4}}, 1, 4, [][]int{[]int{1,2,3,4}}
  T(_t, S(i, r, c), S(matrixReshape(i, r, c)), S(o))

  i, r, c, o = [][]int{[]int{1,2}, []int{3,4}}, 2, 4, [][]int{[]int{1,2}, []int{3,4}}
  T(_t, S(i, r, c), S(matrixReshape(i, r, c)), S(o))
}

func TestMatrixReshapeArray(_t *testing.T) {
  i, r, c, o := [][]int{[]int{1,2}, []int{3,4}}, 1, 4, [][]int{[]int{1,2,3,4}}
  T(_t, S(i, r, c), S(matrixReshapeArray(i, r, c)), S(o))

  i, r, c, o = [][]int{[]int{1,2}, []int{3,4}}, 2, 4, [][]int{[]int{1,2}, []int{3,4}}
  T(_t, S(i, r, c), S(matrixReshapeArray(i, r, c)), S(o))
}
