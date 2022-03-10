package main
import "testing"
import . "main/pkg/testing_helper"

func TestArrangeCoins(_t *testing.T) {
  i, o := 0, 0
  T(_t, S(i), S(arrangeCoins(i)), S(o))

  i, o = 1, 1
  T(_t, S(i), S(arrangeCoins(i)), S(o))

  i, o = 5, 2
  T(_t, S(i), S(arrangeCoins(i)), S(o))

  i, o = 8, 3
  T(_t, S(i), S(arrangeCoins(i)), S(o))
}

func TestArrangeCoinsBit(_t *testing.T) {
  i, o := 0, 0
  T(_t, S(i), S(arrangeCoinsBit(i)), S(o))

  i, o = 1, 1
  T(_t, S(i), S(arrangeCoinsBit(i)), S(o))

  i, o = 5, 2
  T(_t, S(i), S(arrangeCoinsBit(i)), S(o))

  i, o = 8, 3
  T(_t, S(i), S(arrangeCoinsBit(i)), S(o))
}
