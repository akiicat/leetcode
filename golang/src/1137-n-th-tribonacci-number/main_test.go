package main
import "testing"
import . "main/pkg/testing_helper"

func TestTribonacci(_t *testing.T) {
  i, o := 4, 4
  T(_t, S(i), S(tribonacci(i)), S(o))

  i, o = 25, 1389537
  T(_t, S(i), S(tribonacci(i)), S(o))
}

func TestTribonacciRecursive(_t *testing.T) {
  i, o := 4, 4
  T(_t, S(i), S(tribonacciRecursive(i)), S(o))

  i, o = 25, 1389537
  T(_t, S(i), S(tribonacciRecursive(i)), S(o))
}

