package main
import "testing"
import . "main/pkg/testing_helper"

func TestReplaceElements(t *testing.T) {
  i, o := []int{17,18,5,4,6,1}, []int{18,6,6,6,1,-1}
  T(t, S(i), S(replaceElements(i)), S(o))
}

func TestReplaceElementsLeftToRight(t *testing.T) {
  i, o := []int{17,18,5,4,6,1}, []int{18,6,6,6,1,-1}
  T(t, S(i), S(replaceElementsLeftToRight(i)), S(o))
}

