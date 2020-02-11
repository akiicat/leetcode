package main
import "testing"
import . "main/pkg/testing_helper"

func TestAddToArrayForm(t *testing.T) {
  a, k, o := []int{1,2,0,0}, 34, []int{1,2,3,4}
  T(t, S(a, k), S(addToArrayForm(a, k)), S(o))

  a, k, o = []int{2,7,4}, 181, []int{4,5,5}
  T(t, S(a, k), S(addToArrayForm(a, k)), S(o))

  a, k, o = []int{2,1,5}, 806, []int{1,0,2,1}
  T(t, S(a, k), S(addToArrayForm(a, k)), S(o))

  a, k, o = []int{9,9,9,9,9,9,9,9,9,9}, 1, []int{1,0,0,0,0,0,0,0,0,0,0}
  T(t, S(a, k), S(addToArrayForm(a, k)), S(o))

  a, k, o = []int{0}, 23, []int{2,3}
  T(t, S(a, k), S(addToArrayForm(a, k)), S(o))
}

