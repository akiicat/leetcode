package main
import "testing"
import . "main/pkg/testing_helper"

func TestLongestCommonPrefix(_t *testing.T) {
  strs, o := []string{"flower","flow","flight"}, "fl"
  T(_t, S(strs), S(longestCommonPrefix(strs)), S(o))

  strs, o = []string{"a","ac"}, "a"
  T(_t, S(strs), S(longestCommonPrefix(strs)), S(o))

  strs, o = []string{"aa","a"}, "a"
  T(_t, S(strs), S(longestCommonPrefix(strs)), S(o))

  strs, o = []string{"dog","racecar","car"}, ""
  T(_t, S(strs), S(longestCommonPrefix(strs)), S(o))

  strs, o = []string{"","b"}, ""
  T(_t, S(strs), S(longestCommonPrefix(strs)), S(o))

  strs, o = []string{}, ""
  T(_t, S(strs), S(longestCommonPrefix(strs)), S(o))
}

func TestLongestCommonPrefixDivideAndConquer(_t *testing.T) {
  strs, o := []string{"flower","flow","flight"}, "fl"
  T(_t, S(strs), S(longestCommonPrefixDivideAndConquer(strs)), S(o))

  strs, o = []string{"a","ac"}, "a"
  T(_t, S(strs), S(longestCommonPrefixDivideAndConquer(strs)), S(o))

  strs, o = []string{"aa","a"}, "a"
  T(_t, S(strs), S(longestCommonPrefixDivideAndConquer(strs)), S(o))

  strs, o = []string{"dog","racecar","car"}, ""
  T(_t, S(strs), S(longestCommonPrefixDivideAndConquer(strs)), S(o))

  strs, o = []string{"","b"}, ""
  T(_t, S(strs), S(longestCommonPrefixDivideAndConquer(strs)), S(o))

  strs, o = []string{}, ""
  T(_t, S(strs), S(longestCommonPrefixDivideAndConquer(strs)), S(o))
}

func TestLongestCommonPrefixVerticalScanning(_t *testing.T) {
  strs, o := []string{"flower","flow","flight"}, "fl"
  T(_t, S(strs), S(longestCommonPrefixVerticalScanning(strs)), S(o))

  strs, o = []string{"a","ac"}, "a"
  T(_t, S(strs), S(longestCommonPrefixVerticalScanning(strs)), S(o))

  strs, o = []string{"aa","a"}, "a"
  T(_t, S(strs), S(longestCommonPrefixVerticalScanning(strs)), S(o))

  strs, o = []string{"dog","racecar","car"}, ""
  T(_t, S(strs), S(longestCommonPrefixVerticalScanning(strs)), S(o))

  strs, o = []string{"","b"}, ""
  T(_t, S(strs), S(longestCommonPrefixVerticalScanning(strs)), S(o))

  strs, o = []string{}, ""
  T(_t, S(strs), S(longestCommonPrefixVerticalScanning(strs)), S(o))
}

func TestLongestCommonPrefixHorizontalScanning(_t *testing.T) {
  strs, o := []string{"flower","flow","flight"}, "fl"
  T(_t, S(strs), S(longestCommonPrefixHorizontalScanning(strs)), S(o))

  strs, o = []string{"a","ac"}, "a"
  T(_t, S(strs), S(longestCommonPrefixHorizontalScanning(strs)), S(o))

  strs, o = []string{"aa","a"}, "a"
  T(_t, S(strs), S(longestCommonPrefixHorizontalScanning(strs)), S(o))

  strs, o = []string{"dog","racecar","car"}, ""
  T(_t, S(strs), S(longestCommonPrefixHorizontalScanning(strs)), S(o))

  strs, o = []string{"","b"}, ""
  T(_t, S(strs), S(longestCommonPrefixHorizontalScanning(strs)), S(o))

  strs, o = []string{}, ""
  T(_t, S(strs), S(longestCommonPrefixHorizontalScanning(strs)), S(o))
}
