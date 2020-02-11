package main
import "testing"
import . "main/pkg/testing_helper"

func TestGetDecimalValue(_t *testing.T) {
  i, o := "I speak Goat Latin", "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
  T(_t, S(i), S(toGoatLatin(i)), S(o))

  i, o = "The quick brown fox jumped over the lazy dog", "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
  T(_t, S(i), S(toGoatLatin(i)), S(o))
}

