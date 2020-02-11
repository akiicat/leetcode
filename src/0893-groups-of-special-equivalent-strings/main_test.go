package main
import "testing"
import . "main/pkg/testing_helper"

func TestSurfaceArea(t *testing.T) {
  i, o := []string{"abcd","cdab","cbad","xyzz","zzxy","zzyx"}, 3
  T(t, S(i), S(numSpecialEquivGroups(i)), S(o))

  i, o = []string{"abc","acb","bac","bca","cab","cba"}, 3
  T(t, S(i), S(numSpecialEquivGroups(i)), S(o))
}

func TestNumSpecialEquivGroupsStringMap(t *testing.T) {
  i, o := []string{"abcd","cdab","cbad","xyzz","zzxy","zzyx"}, 3
  T(t, S(i), S(numSpecialEquivGroupsStringMap(i)), S(o))

  i, o = []string{"abc","acb","bac","bca","cab","cba"}, 3
  T(t, S(i), S(numSpecialEquivGroupsStringMap(i)), S(o))
}

