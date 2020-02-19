package main
import "testing"
import . "main/pkg/testing_helper"

func TestRepeatedNTimes(_t *testing.T) {
  i, o := []int{1,2,3,3}, 3
  T(_t, S(i), S(repeatedNTimes(i)), S(o))

  i, o = []int{2,1,2,5,3,2}, 2
  T(_t, S(i), S(repeatedNTimes(i)), S(o))

  i, o = []int{5,1,5,2,5,3,5,4}, 5
  T(_t, S(i), S(repeatedNTimes(i)), S(o))

  i, o = []int{2,6,2,1}, 2
  T(_t, S(i), S(repeatedNTimes(i)), S(o))
}

func TestRepeatedNTimesArray(_t *testing.T) {
  i, o := []int{1,2,3,3}, 3
  T(_t, S(i), S(repeatedNTimesArray(i)), S(o))

  i, o = []int{2,1,2,5,3,2}, 2
  T(_t, S(i), S(repeatedNTimesArray(i)), S(o))

  i, o = []int{5,1,5,2,5,3,5,4}, 5
  T(_t, S(i), S(repeatedNTimesArray(i)), S(o))

  i, o = []int{2,6,2,1}, 2
  T(_t, S(i), S(repeatedNTimesArray(i)), S(o))
}

func TestRepeatedNTimesHash(_t *testing.T) {
  i, o := []int{1,2,3,3}, 3
  T(_t, S(i), S(repeatedNTimesHash(i)), S(o))

  i, o = []int{2,1,2,5,3,2}, 2
  T(_t, S(i), S(repeatedNTimesHash(i)), S(o))

  i, o = []int{5,1,5,2,5,3,5,4}, 5
  T(_t, S(i), S(repeatedNTimesHash(i)), S(o))

  i, o = []int{2,6,2,1}, 2
  T(_t, S(i), S(repeatedNTimesHash(i)), S(o))
}

