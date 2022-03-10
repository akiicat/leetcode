package main
import "testing"
import . "main/pkg/testing_helper"

func TestCountPrimes(_t *testing.T) {
  i, o := 10, 4
  T(_t, S(i), S(countPrimes(i)), S(o))

  i, o = 20, 8
  T(_t, S(i), S(countPrimes(i)), S(o))
}

func TestCountPrimesIsPrime(_t *testing.T) {
  i, o := 10, 4
  T(_t, S(i), S(countPrimesIsPrime(i)), S(o))

  i, o = 20, 8
  T(_t, S(i), S(countPrimesIsPrime(i)), S(o))
}
