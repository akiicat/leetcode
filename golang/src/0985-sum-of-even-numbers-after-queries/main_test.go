package main
import "testing"
import . "main/pkg/testing_helper"

func TestSumEvenAfterQueries(t *testing.T) {
  i, queries, o := []int{1,2,3,4}, [][]int{[]int{1,0},[]int{-3,1},[]int{-4,0},[]int{2,3}}, []int{8,6,2,4}
  T(t, S(i, queries), S(sumEvenAfterQueries(i, queries)), S(o))
}

