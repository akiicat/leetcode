package main
import "testing"
import . "main/pkg/testing_helper"
import "unicode"

func TestReformat(_t *testing.T) {
  i, o := "a0b1c2", "0a1b2c"
  T(_t, S(i), S(Valid(reformat(i), o)), S(o))

  i, o = "leetcode", ""
  T(_t, S(i), S(Valid(reformat(i), o)), S(o))

  i, o = "1229857369", ""
  T(_t, S(i), S(Valid(reformat(i), o)), S(o))

  i, o = "covid2019", "c2o0v1i9d"
  T(_t, S(i), S(Valid(reformat(i), o)), S(o))

  i, o = "ab123", "1a2b3"
  T(_t, S(i), S(Valid(reformat(i), o)), S(o))
}

func TestReformatSplit(_t *testing.T) {
  i, o := "a0b1c2", "0a1b2c"
  T(_t, S(i), S(Valid(reformatSplit(i), o)), S(o))

  i, o = "leetcode", ""
  T(_t, S(i), S(Valid(reformatSplit(i), o)), S(o))

  i, o = "1229857369", ""
  T(_t, S(i), S(Valid(reformatSplit(i), o)), S(o))

  i, o = "covid2019", "c2o0v1i9d"
  T(_t, S(i), S(Valid(reformatSplit(i), o)), S(o))

  i, o = "ab123", "1a2b3"
  T(_t, S(i), S(Valid(reformatSplit(i), o)), S(o))
}

// return r if invalid
// return o if valid
func Valid(r, o string) string {
  if len(r) != len(o) {
    return r
  }

  if len(r) == 0 {
    return o
  }

  m := map[byte]int{}
  m[r[0]]++
  m[o[0]]--

  for i := 0; i < len(r)-1; i++ {
    if unicode.IsLower(rune(r[i])) != unicode.IsDigit(rune(r[i+1])) {
      return r
    }

    m[r[i+1]]++
    m[o[i+1]]--
  }

  for _, v := range m {
    if v != 0 {
      return r
    }
  }

  return o
}

