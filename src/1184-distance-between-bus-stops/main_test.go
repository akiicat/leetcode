package main
import "testing"
import . "main/pkg/testing_helper"

func TestDistanceBetweenBusStops(_t *testing.T) {
  i, s, t, o := []int{1,2,3,4}, 0, 1, 1
  T(_t, S(i, s, t), S(distanceBetweenBusStops(i, s, t)), S(o))

  i, s, t, o = []int{1,2,3,4}, 0, 2, 3
  T(_t, S(i, s, t), S(distanceBetweenBusStops(i, s, t)), S(o))

  i, s, t, o = []int{1,2,3,4}, 0, 3, 4
  T(_t, S(i, s, t), S(distanceBetweenBusStops(i, s, t)), S(o))
}

