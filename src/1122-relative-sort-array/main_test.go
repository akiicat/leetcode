package main
import "testing"
import . "main/pkg/testing_helper"

func TestRelativeSortArray(t *testing.T) {
  i1, i2, o := []int{2,3,1,3,2,4,6,7,9,2,19}, []int{2,1,4,3,9,6}, []int{2,2,2,1,4,3,3,9,6,7,19}
  T(t, S(i1, i2), S(relativeSortArray(i1, i2)), S(o))

  i1, i2, o = []int{28,6,22,8,44,17}, []int{22,28,8,6}, []int{22,28,8,6,17,44}
  T(t, S(i1, i2), S(relativeSortArray(i1, i2)), S(o))
}

func TestRelativeSortArrayMergeSort(t *testing.T) {
  i1, i2, o := []int{2,3,1,3,2,4,6,7,9,2,19}, []int{2,1,4,3,9,6}, []int{2,2,2,1,4,3,3,9,6,7,19}
  T(t, S(i1, i2), S(relativeSortArrayMergeSort(i1, i2)), S(o))

  i1, i2, o = []int{28,6,22,8,44,17}, []int{22,28,8,6}, []int{22,28,8,6,17,44}
  T(t, S(i1, i2), S(relativeSortArrayMergeSort(i1, i2)), S(o))
}

