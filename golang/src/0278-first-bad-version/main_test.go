package main
import "testing"
import . "main/pkg/testing_helper"

func TestFirstBadVersion(_t *testing.T) {
  i, o := 5, 4
  setBadVersion(o)
  T(_t, S(i), S(firstBadVersion(i)), S(o))

  i, o = 1, 1
  setBadVersion(o)
  T(_t, S(i), S(firstBadVersion(i)), S(o))
}

