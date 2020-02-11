package main
import "testing"
import . "main/pkg/testing_helper"

func TestSearch(_t *testing.T) {
  i, t, o := []int{-1,0,3,5,9,12}, 9, 4
  T(_t, S(i, t), S(search(i, t)), S(o))

  i, t, o = []int{-1,0,3,5,9,12}, 2, -1
  T(_t, S(i, t), S(search(i, t)), S(o))
}

func TestSearchRecursive(_t *testing.T) {
  i, t, o := []int{-1,0,3,5,9,12}, 9, 4
  T(_t, S(i, t), S(searchRecursive(i, t)), S(o))

  i, t, o = []int{-1,0,3,5,9,12}, 2, -1
  T(_t, S(i, t), S(searchRecursive(i, t)), S(o))
}

