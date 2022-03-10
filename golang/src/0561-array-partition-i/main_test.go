package main
import "testing"
import . "main/pkg/testing_helper"

func TestArrayPairSum(_t *testing.T) {
  i, o := []int{1,4,3,2}, 4
  T(_t, S(i), S(arrayPairSum(i)), S(o))
}

func TestArrayPairSumSort(_t *testing.T) {
  i, o := []int{1,4,3,2}, 4
  T(_t, S(i), S(arrayPairSumSort(i)), S(o))
}

