package main
import "testing"
import . "main/pkg/testing_helper"

func TestSortedSquares(t *testing.T) {
  i, o := []int{-4,-1,0,3,10}, []int{0,1,9,16,100}
  T(t, S(i), S(sortedSquares(i)), S(o))

  i, o = []int{-7,-3,2,3,11}, []int{4,9,9,49,121}
  T(t, S(i), S(sortedSquares(i)), S(o))
}

func TestSortedSquaresReverse(t *testing.T) {
  i, o := []int{-4,-1,0,3,10}, []int{0,1,9,16,100}
  T(t, S(i), S(sortedSquaresReverse(i)), S(o))

  i, o = []int{-7,-3,2,3,11}, []int{4,9,9,49,121}
  T(t, S(i), S(sortedSquaresReverse(i)), S(o))
}

