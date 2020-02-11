package main
import "testing"
import . "main/pkg/testing_helper"

func TestBackspaceCompare(_t *testing.T) {
  s, t, o := "ab#c", "ad#c", true
  T(_t, S(s, t), S(backspaceCompare(s, t)), S(o))

  s, t, o = "ab##", "c#d#", true
  T(_t, S(s, t), S(backspaceCompare(s, t)), S(o))

  s, t, o = "a##c", "#a#c", true
  T(_t, S(s, t), S(backspaceCompare(s, t)), S(o))

  s, t, o = "a#c", "b", false
  T(_t, S(s, t), S(backspaceCompare(s, t)), S(o))

  s, t, o = "y#fo##f", "y#f#o##f", true
  T(_t, S(s, t), S(backspaceCompare(s, t)), S(o))
}

func TestBackspaceCompareTwoPointer(_t *testing.T) {
  s, t, o := "ab#c", "ad#c", true
  T(_t, S(s, t), S(backspaceCompareTwoPointer(s, t)), S(o))

  s, t, o = "ab##", "c#d#", true
  T(_t, S(s, t), S(backspaceCompareTwoPointer(s, t)), S(o))

  s, t, o = "a##c", "#a#c", true
  T(_t, S(s, t), S(backspaceCompareTwoPointer(s, t)), S(o))

  s, t, o = "a#c", "b", false
  T(_t, S(s, t), S(backspaceCompareTwoPointer(s, t)), S(o))

  s, t, o = "y#fo##f", "y#f#o##f", true
  T(_t, S(s, t), S(backspaceCompareTwoPointer(s, t)), S(o))
}

func TestBackspaceCompareStack(_t *testing.T) {
  s, t, o := "ab#c", "ad#c", true
  T(_t, S(s, t), S(backspaceCompareStack(s, t)), S(o))

  s, t, o = "ab##", "c#d#", true
  T(_t, S(s, t), S(backspaceCompareStack(s, t)), S(o))

  s, t, o = "a##c", "#a#c", true
  T(_t, S(s, t), S(backspaceCompareStack(s, t)), S(o))

  s, t, o = "a#c", "b", false
  T(_t, S(s, t), S(backspaceCompareStack(s, t)), S(o))

  s, t, o = "y#fo##f", "y#f#o##f", true
  T(_t, S(s, t), S(backspaceCompareStack(s, t)), S(o))
}

