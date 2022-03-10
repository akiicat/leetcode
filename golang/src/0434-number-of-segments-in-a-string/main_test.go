package main
import "testing"
import . "main/pkg/testing_helper"

func TestCountSegments(_t *testing.T) {
  i, o := "Hello, my name is John", 5
  T(_t, S(i), S(countSegments(i)), S(o))

  i, o = "Hello", 1
  T(_t, S(i), S(countSegments(i)), S(o))

  i, o = "", 0
  T(_t, S(i), S(countSegments(i)), S(o))

  i, o = "           ", 0
  T(_t, S(i), S(countSegments(i)), S(o))
}

func TestCountSegmentsCountWords(_t *testing.T) {
  i, o := "Hello, my name is John", 5
  T(_t, S(i), S(countSegmentsCountWords(i)), S(o))

  i, o = "Hello", 1
  T(_t, S(i), S(countSegmentsCountWords(i)), S(o))

  i, o = "", 0
  T(_t, S(i), S(countSegmentsCountWords(i)), S(o))

  i, o = "           ", 0
  T(_t, S(i), S(countSegmentsCountWords(i)), S(o))
}

