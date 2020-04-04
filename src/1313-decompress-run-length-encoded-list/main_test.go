package main
import "testing"
import . "main/pkg/testing_helper"

func TestDecompressRLElist(t *testing.T) {
  i, o := []int{1,2,3,4}, []int{2,4,4,4}
  T(t, S(i), S(decompressRLElist(i)), S(o))

  i, o = []int{1,1,2,3}, []int{1,3,3}
  T(t, S(i), S(decompressRLElist(i)), S(o))
}
