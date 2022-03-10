package main
import "testing"
import . "main/pkg/testing_helper"

func TestFreqAlphabets(_t *testing.T) {
  i, o := "10#11#12", "jkab"
  T(_t, S(i), S(freqAlphabets(i)), S(o))

  i, o = "1326#", "acz"
  T(_t, S(i), S(freqAlphabets(i)), S(o))

  i, o = "25#", "y"
  T(_t, S(i), S(freqAlphabets(i)), S(o))

  i, o = "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#", "abcdefghijklmnopqrstuvwxyz"
  T(_t, S(i), S(freqAlphabets(i)), S(o))
}

func TestFreqAlphabetsBackward(_t *testing.T) {
  i, o := "10#11#12", "jkab"
  T(_t, S(i), S(freqAlphabetsBackward(i)), S(o))

  i, o = "1326#", "acz"
  T(_t, S(i), S(freqAlphabetsBackward(i)), S(o))

  i, o = "25#", "y"
  T(_t, S(i), S(freqAlphabetsBackward(i)), S(o))

  i, o = "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#", "abcdefghijklmnopqrstuvwxyz"
  T(_t, S(i), S(freqAlphabetsBackward(i)), S(o))
}

