package main
import "testing"
import . "main/pkg/testing_helper"

func TestHasAlternatingBits(_t *testing.T) {
  i, o :=  0, true
  T(_t, S(i), S(hasAlternatingBits(i)), S(o))

  i, o =  1, true
  T(_t, S(i), S(hasAlternatingBits(i)), S(o))

  i, o =  2, true
  T(_t, S(i), S(hasAlternatingBits(i)), S(o))

  i, o =  3, false
  T(_t, S(i), S(hasAlternatingBits(i)), S(o))

  i, o =  4, false
  T(_t, S(i), S(hasAlternatingBits(i)), S(o))

  i, o =  5, true
  T(_t, S(i), S(hasAlternatingBits(i)), S(o))

  i, o =  1431655765, true
  T(_t, S(i), S(hasAlternatingBits(i)), S(o))
}

func TestHasAlternatingBitsCancelBits(_t *testing.T) {
  i, o :=  0, true
  T(_t, S(i), S(hasAlternatingBitsCancelBits(i)), S(o))

  i, o =  1, true
  T(_t, S(i), S(hasAlternatingBitsCancelBits(i)), S(o))

  i, o =  2, true
  T(_t, S(i), S(hasAlternatingBitsCancelBits(i)), S(o))

  i, o =  3, false
  T(_t, S(i), S(hasAlternatingBitsCancelBits(i)), S(o))

  i, o =  4, false
  T(_t, S(i), S(hasAlternatingBitsCancelBits(i)), S(o))

  i, o =  5, true
  T(_t, S(i), S(hasAlternatingBitsCancelBits(i)), S(o))

  i, o =  1431655765, true
  T(_t, S(i), S(hasAlternatingBitsCancelBits(i)), S(o))
}

func TestHasAlternatingBitsCompleteBits(_t *testing.T) {
  i, o :=  0, true
  T(_t, S(i), S(hasAlternatingBitsCompleteBits(i)), S(o))

  i, o =  1, true
  T(_t, S(i), S(hasAlternatingBitsCompleteBits(i)), S(o))

  i, o =  2, true
  T(_t, S(i), S(hasAlternatingBitsCompleteBits(i)), S(o))

  i, o =  3, false
  T(_t, S(i), S(hasAlternatingBitsCompleteBits(i)), S(o))

  i, o =  4, false
  T(_t, S(i), S(hasAlternatingBitsCompleteBits(i)), S(o))

  i, o =  5, true
  T(_t, S(i), S(hasAlternatingBitsCompleteBits(i)), S(o))

  i, o =  1431655765, true
  T(_t, S(i), S(hasAlternatingBitsCompleteBits(i)), S(o))
}

