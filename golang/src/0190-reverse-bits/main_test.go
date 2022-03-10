package main
import "testing"
import . "main/pkg/testing_helper"

func TestReverseBits(_t *testing.T) {
  //  43261596 (00000010100101000001111010011100)
  //  964176192 (00111001011110000010100101000000)
  i, o := uint32(43261596), uint32(964176192)
  T(_t, S(i), S(reverseBits(i)), S(o))
}
