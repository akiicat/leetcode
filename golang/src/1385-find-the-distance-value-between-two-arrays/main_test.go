package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindTheDistanceValue(_t *testing.T) {
  arr1, arr2, d, o := []int{4,5,8}, []int{10,9,1,8}, 2, 2
  T(_t, S(arr1, arr2, d), S(findTheDistanceValue(arr1, arr2, d)), S(o))

  arr1, arr2, d, o = []int{1,4,2,3}, []int{-4,-3,6,10,20,30}, 3, 2
  T(_t, S(arr1, arr2, d), S(findTheDistanceValue(arr1, arr2, d)), S(o))

  arr1, arr2, d, o = []int{2,1,100,3}, []int{-5,-2,10,-3,7}, 6, 1
  T(_t, S(arr1, arr2, d), S(findTheDistanceValue(arr1, arr2, d)), S(o))
}

