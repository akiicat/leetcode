package main
import "testing"
import . "main/pkg/testing_helper"

func TestLastStoneWeight(_t *testing.T) {
  i, o := []int{2,7,4,1,8,1}, 1
  T(_t, S(i), S(lastStoneWeight(i)), S(o))

  i, o = []int{7,6,7,6,9}, 3
  T(_t, S(i), S(lastStoneWeight(i)), S(o))

  i, o = []int{10,4,2,10}, 2
  T(_t, S(i), S(lastStoneWeight(i)), S(o))

  i, o = []int{9,3,2,10}, 0
  T(_t, S(i), S(lastStoneWeight(i)), S(o))

  i, o = []int{1,3}, 2
  T(_t, S(i), S(lastStoneWeight(i)), S(o))

  i, o = []int{1}, 1
  T(_t, S(i), S(lastStoneWeight(i)), S(o))
}

func TestLastStoneWeightSort(_t *testing.T) {
  i, o := []int{2,7,4,1,8,1}, 1
  T(_t, S(i), S(lastStoneWeightSort(i)), S(o))

  i, o = []int{7,6,7,6,9}, 3
  T(_t, S(i), S(lastStoneWeightSort(i)), S(o))

  i, o = []int{10,4,2,10}, 2
  T(_t, S(i), S(lastStoneWeightSort(i)), S(o))

  i, o = []int{9,3,2,10}, 0
  T(_t, S(i), S(lastStoneWeightSort(i)), S(o))

  i, o = []int{1,3}, 2
  T(_t, S(i), S(lastStoneWeightSort(i)), S(o))

  i, o = []int{1}, 1
  T(_t, S(i), S(lastStoneWeightSort(i)), S(o))
}

