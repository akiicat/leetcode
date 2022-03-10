package main
import "testing"
import . "main/pkg/testing_helper"

func TestNumEquivDominoPairs(t *testing.T) {
  i, o := [][]int{[]int{1,2},[]int{2,1},[]int{3,4},[]int{5,6}}, 1
  T(t, S(i), S(numEquivDominoPairs(i)), S(o))
}

