package main
import "testing"
import . "main/pkg/testing_helper"

func TestCanPlaceFlowers(_t *testing.T) {
  i, n, o := []int{1,0,0,0,1}, 1, true
  T(_t, S(i, n), S(canPlaceFlowers(i, n)), S(o))

  i, n, o = []int{1,0,0,0,1}, 2, false
  T(_t, S(i, n), S(canPlaceFlowers(i, n)), S(o))

  i, n, o = []int{1,0,0,0,0,1}, 2, false
  T(_t, S(i, n), S(canPlaceFlowers(i, n)), S(o))

  i, n, o = []int{1,0,0,0,1,0,0}, 2, true
  T(_t, S(i, n), S(canPlaceFlowers(i, n)), S(o))

  i, n, o = []int{0,0,1,0,0,0,1}, 2, true
  T(_t, S(i, n), S(canPlaceFlowers(i, n)), S(o))

  i, n, o = []int{0,1,0}, 1, false
  T(_t, S(i, n), S(canPlaceFlowers(i, n)), S(o))
}
