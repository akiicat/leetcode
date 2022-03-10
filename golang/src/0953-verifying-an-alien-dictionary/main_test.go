package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsAlienSorted(_t *testing.T) {
  i, order, o := []string{"hello","leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true
  T(_t, S(i, order), S(isAlienSorted(i, order)), S(o))

  i, order, o = []string{"word","world","row"}, "worldabcefghijkmnpqstuvxyz", false
  T(_t, S(i, order), S(isAlienSorted(i, order)), S(o))

  i, order, o = []string{"apple","app"}, "abcdefghijklmnopqrstuvwxyz", false
  T(_t, S(i, order), S(isAlienSorted(i, order)), S(o))

  i, order, o = []string{"fxasxpc","dfbdrifhp","nwzgs","cmwqriv","ebulyfyve","miracx","sxckdwzv","dtijzluhts","wwbmnge","qmjwymmyox"}, "zkgwaverfimqxbnctdplsjyohu", false
  T(_t, S(i, order), S(isAlienSorted(i, order)), S(o))
}

