package main
import "testing"
import . "main/pkg/testing_helper"

func TestArrayRankTransform(t *testing.T) {
  i, o := []int{40,10,20,30}, []int{4,1,2,3}
  T(t, S(i), S(arrayRankTransform(i)), S(o))

  i, o = []int{100,100,100}, []int{1,1,1}
  T(t, S(i), S(arrayRankTransform(i)), S(o))

  i, o = []int{37,12,28,9,100,56,80,5,12}, []int{5,3,4,2,8,6,7,1,3}
  T(t, S(i), S(arrayRankTransform(i)), S(o))
}

