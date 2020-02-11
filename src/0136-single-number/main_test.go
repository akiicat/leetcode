package main
import "testing"
import . "main/pkg/testing_helper"

func TestSingleNumber(_t *testing.T) {
  i, o := []int{2,2,1}, 1
  T(_t, S(i), S(singleNumber(i)), S(o))

  i, o = []int{4,1,2,1,2}, 4
  T(_t, S(i), S(singleNumber(i)), S(o))
}

func TestSingleNumberHash(_t *testing.T) {
  i, o := []int{2,2,1}, 1
  T(_t, S(i), S(singleNumberHash(i)), S(o))

  i, o = []int{4,1,2,1,2}, 4
  T(_t, S(i), S(singleNumberHash(i)), S(o))
}

func TestSingleNumberBruteForce(_t *testing.T) {
  i, o := []int{2,2,1}, 1
  T(_t, S(i), S(singleNumberBruteForce(i)), S(o))

  i, o = []int{4,1,2,1,2}, 4
  T(_t, S(i), S(singleNumberBruteForce(i)), S(o))
}

