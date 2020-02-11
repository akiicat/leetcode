package main
import "testing"
import . "main/pkg/testing_helper"

func TestPeakIndexInMountainArray(_t *testing.T) {
  i, o := []int{0,1,0}, 1
  T(_t, S(i), S(peakIndexInMountainArray(i)), S(o))

  i, o = []int{0,2,1,0}, 1
  T(_t, S(i), S(peakIndexInMountainArray(i)), S(o))

  i, o = []int{40,48,61,75,100,99,98,39,30,10}, 4
  T(_t, S(i), S(peakIndexInMountainArray(i)), S(o))
}

func TestPeakIndexInMountainArrayBinarySearch(_t *testing.T) {
  i, o := []int{0,1,0}, 1
  T(_t, S(i), S(peakIndexInMountainArrayBinarySearch(i)), S(o))

  i, o = []int{0,2,1,0}, 1
  T(_t, S(i), S(peakIndexInMountainArrayBinarySearch(i)), S(o))

  i, o = []int{40,48,61,75,100,99,98,39,30,10}, 4
  T(_t, S(i), S(peakIndexInMountainArrayBinarySearch(i)), S(o))
}

func TestPeakIndexInMountainArrayBinarySearchRecursive(_t *testing.T) {
  i, o := []int{0,1,0}, 1
  T(_t, S(i), S(peakIndexInMountainArrayBinarySearchRecursive(i)), S(o))

  i, o = []int{0,2,1,0}, 1
  T(_t, S(i), S(peakIndexInMountainArrayBinarySearchRecursive(i)), S(o))

  i, o = []int{40,48,61,75,100,99,98,39,30,10}, 4
  T(_t, S(i), S(peakIndexInMountainArrayBinarySearchRecursive(i)), S(o))
}

