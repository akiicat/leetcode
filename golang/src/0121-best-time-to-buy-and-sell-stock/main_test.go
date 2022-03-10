package main
import "testing"
import . "main/pkg/testing_helper"

func TestMaxProfit(_t *testing.T) {
  i, o := []int{7,1,5,3,6,4}, 5
  T(_t, S(i), S(maxProfit(i)), S(o))

  i, o = []int{7,6,4,3,1}, 0
  T(_t, S(i), S(maxProfit(i)), S(o))

  i, o = []int{}, 0
  T(_t, S(i), S(maxProfit(i)), S(o))
}

func TestMaxProfitSubArray(_t *testing.T) {
  i, o := []int{7,1,5,3,6,4}, 5
  T(_t, S(i), S(maxProfitSubArray(i)), S(o))

  i, o = []int{7,6,4,3,1}, 0
  T(_t, S(i), S(maxProfitSubArray(i)), S(o))

  i, o = []int{}, 0
  T(_t, S(i), S(maxProfitSubArray(i)), S(o))
}
