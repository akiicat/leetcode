package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindUnsortedSubarray(_t *testing.T) {
  i, o := []int{2,6,4,8,10,9,15}, 5
  T(_t, S(i), S(findUnsortedSubarray(i)), S(o))

  i, o = []int{1}, 0
  T(_t, S(i), S(findUnsortedSubarray(i)), S(o))

  i, o = []int{1,3,2,2,2}, 4
  T(_t, S(i), S(findUnsortedSubarray(i)), S(o))

  i, o = []int{1,2,3,3,3}, 0
  T(_t, S(i), S(findUnsortedSubarray(i)), S(o))

  i, o = []int{1,3,2,3,3}, 2
  T(_t, S(i), S(findUnsortedSubarray(i)), S(o))
}

func TestFindUnsortedSubarrayStack(_t *testing.T) {
  i, o := []int{2,6,4,8,10,9,15}, 5
  T(_t, S(i), S(findUnsortedSubarrayStack(i)), S(o))

  i, o = []int{1}, 0
  T(_t, S(i), S(findUnsortedSubarrayStack(i)), S(o))

  i, o = []int{1,3,2,2,2}, 4
  T(_t, S(i), S(findUnsortedSubarrayStack(i)), S(o))

  i, o = []int{1,2,3,3,3}, 0
  T(_t, S(i), S(findUnsortedSubarrayStack(i)), S(o))

  i, o = []int{1,3,2,3,3}, 2
  T(_t, S(i), S(findUnsortedSubarrayStack(i)), S(o))
}

func TestFindUnsortedSubarraySort(_t *testing.T) {
  i, o := []int{2,6,4,8,10,9,15}, 5
  T(_t, S(i), S(findUnsortedSubarraySort(i)), S(o))

  i, o = []int{1}, 0
  T(_t, S(i), S(findUnsortedSubarraySort(i)), S(o))

  i, o = []int{1,3,2,2,2}, 4
  T(_t, S(i), S(findUnsortedSubarraySort(i)), S(o))

  i, o = []int{1,2,3,3,3}, 0
  T(_t, S(i), S(findUnsortedSubarraySort(i)), S(o))

  i, o = []int{1,3,2,3,3}, 2
  T(_t, S(i), S(findUnsortedSubarraySort(i)), S(o))
}

func TestFindUnsortedSubarrayBruteForce(_t *testing.T) {
  i, o := []int{2,6,4,8,10,9,15}, 5
  T(_t, S(i), S(findUnsortedSubarrayBruteForce(i)), S(o))

  i, o = []int{1}, 0
  T(_t, S(i), S(findUnsortedSubarrayBruteForce(i)), S(o))

  i, o = []int{1,3,2,2,2}, 4
  T(_t, S(i), S(findUnsortedSubarrayBruteForce(i)), S(o))

  i, o = []int{1,2,3,3,3}, 0
  T(_t, S(i), S(findUnsortedSubarrayBruteForce(i)), S(o))

  i, o = []int{1,3,2,3,3}, 2
  T(_t, S(i), S(findUnsortedSubarrayBruteForce(i)), S(o))
}

