package main
import "testing"
import . "main/pkg/testing_helper"

func TestQueryString(_t *testing.T) {
  s, n, o := "0110", 3, true
  T(_t, S(s, n), S(queryString(s, n)), S(o))

  s, n, o = "0110", 4, false
  T(_t, S(s, n), S(queryString(s, n)), S(o))
}

