package main
import "testing"
import . "main/pkg/testing_helper"

func TestDistributeCandies(t *testing.T) {
  c, n, o := 7, 4, []int{1,2,3,1}
  T(t, S(c, n), S(distributeCandies(c, n)), S(o))

  c, n, o = 10, 3, []int{5,2,3}
  T(t, S(c, n), S(distributeCandies(c, n)), S(o))

  c, n, o = 12, 3, []int{5,4,3}
  T(t, S(c, n), S(distributeCandies(c, n)), S(o))

  c, n, o = 21, 3, []int{5,7,9}
  T(t, S(c, n), S(distributeCandies(c, n)), S(o))

  c, n, o = 23, 3, []int{7,7,9}
  T(t, S(c, n), S(distributeCandies(c, n)), S(o))

  c, n, o = 60, 4, []int{15,18,15,12}
  T(t, S(c, n), S(distributeCandies(c, n)), S(o))

  c, n, o = 90, 4, []int{27,18,21,24}
  T(t, S(c, n), S(distributeCandies(c, n)), S(o))
}

func TestDistributeCandiesMath(t *testing.T) {
  c, n, o := 7, 4, []int{1,2,3,1}
  T(t, S(c, n), S(distributeCandiesMath(c, n)), S(o))

  c, n, o = 10, 3, []int{5,2,3}
  T(t, S(c, n), S(distributeCandiesMath(c, n)), S(o))

  c, n, o = 12, 3, []int{5,4,3}
  T(t, S(c, n), S(distributeCandiesMath(c, n)), S(o))

  c, n, o = 21, 3, []int{5,7,9}
  T(t, S(c, n), S(distributeCandiesMath(c, n)), S(o))

  c, n, o = 23, 3, []int{7,7,9}
  T(t, S(c, n), S(distributeCandiesMath(c, n)), S(o))

  c, n, o = 60, 4, []int{15,18,15,12}
  T(t, S(c, n), S(distributeCandiesMath(c, n)), S(o))

  c, n, o = 90, 4, []int{27,18,21,24}
  T(t, S(c, n), S(distributeCandiesMath(c, n)), S(o))
}

