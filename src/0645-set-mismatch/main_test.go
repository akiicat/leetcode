package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindErrorNums(_t *testing.T) {
  i, o := []int{1,2,2,4}, []int{2,3}
  T(_t, S(i), S(findErrorNums(i)), S(o))

  i, o = []int{4,2,2,1}, []int{2,3}
  T(_t, S(i), S(findErrorNums(i)), S(o))

  i, o = []int{2,2}, []int{2,1}
  T(_t, S(i), S(findErrorNums(i)), S(o))
}

func TestFindErrorNumsMath(_t *testing.T) {
  i, o := []int{1,2,2,4}, []int{2,3}
  T(_t, S(i), S(findErrorNumsMath(i)), S(o))

  i, o = []int{4,2,2,1}, []int{2,3}
  T(_t, S(i), S(findErrorNumsMath(i)), S(o))

  i, o = []int{2,2}, []int{2,1}
  T(_t, S(i), S(findErrorNumsMath(i)), S(o))
}

func TestFindErrorNumsNegative(_t *testing.T) {
  i, o := []int{1,2,2,4}, []int{2,3}
  T(_t, S(i), S(findErrorNumsNegative(i)), S(o))

  i, o = []int{4,2,2,1}, []int{2,3}
  T(_t, S(i), S(findErrorNumsNegative(i)), S(o))

  i, o = []int{2,2}, []int{2,1}
  T(_t, S(i), S(findErrorNumsNegative(i)), S(o))
}

func TestFindErrorNumsArray(_t *testing.T) {
  i, o := []int{1,2,2,4}, []int{2,3}
  T(_t, S(i), S(findErrorNumsArray(i)), S(o))

  i, o = []int{4,2,2,1}, []int{2,3}
  T(_t, S(i), S(findErrorNumsArray(i)), S(o))

  i, o = []int{2,2}, []int{2,1}
  T(_t, S(i), S(findErrorNumsArray(i)), S(o))
}

