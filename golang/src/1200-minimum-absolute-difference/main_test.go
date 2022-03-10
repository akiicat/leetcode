package main
import "testing"
import . "main/pkg/testing_helper"

func TestMaxNumberOfBalloons(t *testing.T) {
  i, o := []int{4,2,1,3}, [][]int{[]int{1,2}, []int{2,3}, []int{3,4}}
  T(t, S(i), S(minimumAbsDifference(i)), S(o))

  i, o = []int{1,3,6,10,15}, [][]int{[]int{1,3}}
  T(t, S(i), S(minimumAbsDifference(i)), S(o))

  i, o = []int{3,8,-10,23,19,-4,-14,27}, [][]int{[]int{-14,-10}, []int{19,23}, []int{23,27}}
  T(t, S(i), S(minimumAbsDifference(i)), S(o))
}

func TestMinimumAbsDifferenceTwoLoop(t *testing.T) {
  i, o := []int{4,2,1,3}, [][]int{[]int{1,2}, []int{2,3}, []int{3,4}}
  T(t, S(i), S(minimumAbsDifferenceTwoLoop(i)), S(o))

  i, o = []int{1,3,6,10,15}, [][]int{[]int{1,3}}
  T(t, S(i), S(minimumAbsDifferenceTwoLoop(i)), S(o))

  i, o = []int{3,8,-10,23,19,-4,-14,27}, [][]int{[]int{-14,-10}, []int{19,23}, []int{23,27}}
  T(t, S(i), S(minimumAbsDifferenceTwoLoop(i)), S(o))
}

