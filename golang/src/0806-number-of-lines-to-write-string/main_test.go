package main
import "testing"
import . "main/pkg/testing_helper"

func TestNumberOfLines(_t *testing.T) {
  widths, s, o := []int{10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10}, "abcdefghijklmnopqrstuvwxyz", []int{3,60}
  T(_t, S(widths, s), S(numberOfLines(widths, s)), S(o))
}

