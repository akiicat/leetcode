package main
import "testing"
import . "main/pkg/testing_helper"

func TestShortestToChar(_t *testing.T) {
  s, c, o := "loveleetcode", byte('e'), []int{3,2,1,0,1,0,0,1,2,2,1,0}
  T(_t, S(s, c), S(shortestToChar(s, c)), S(o))
}

func TestShortestToCharTwoPass(_t *testing.T) {
  s, c, o := "loveleetcode", byte('e'), []int{3,2,1,0,1,0,0,1,2,2,1,0}
  T(_t, S(s, c), S(shortestToCharTwoPass(s, c)), S(o))
}

