package main
import "testing"
import . "main/pkg/testing_helper"

func TestFairCandySwap(t *testing.T) {
  a, b, o := []int{1,1}, []int{2,2}, []int{1,2}
  T(t, S(a, b), S(fairCandySwap(a, b)), S(o))

  a, b, o = []int{1,2}, []int{2,3}, []int{1,2}
  T(t, S(a, b), S(fairCandySwap(a, b)), S(o))

  a, b, o = []int{2}, []int{1,3}, []int{2,3}
  T(t, S(a, b), S(fairCandySwap(a, b)), S(o))

  a, b, o = []int{1,2,5}, []int{2,4}, []int{5,4}
  T(t, S(a, b), S(fairCandySwap(a, b)), S(o))

  a, b, o = []int{35,17,4,24,10}, []int{63,21}, []int{24,21}
  T(t, S(a, b), S(fairCandySwap(a, b)), S(o))
}

func TestFairCandySwapSort(t *testing.T) {
  a, b, o := []int{1,1}, []int{2,2}, []int{1,2}
  T(t, S(a, b), S(fairCandySwapSort(a, b)), S(o))

  a, b, o = []int{1,2}, []int{2,3}, []int{1,2}
  T(t, S(a, b), S(fairCandySwapSort(a, b)), S(o))

  a, b, o = []int{2}, []int{1,3}, []int{2,3}
  T(t, S(a, b), S(fairCandySwapSort(a, b)), S(o))

  a, b, o = []int{1,2,5}, []int{2,4}, []int{5,4}
  T(t, S(a, b), S(fairCandySwapSort(a, b)), S(o))

  a, b, o = []int{35,17,4,24,10}, []int{63,21}, []int{24,21}
  T(t, S(a, b), S(fairCandySwapSort(a, b)), S(o))
}

