package main
import "strconv"

// T: O(n)
// M: O(1)
// -- start --

// Peek
// T: O(n)
// M: O(1)
func freqAlphabets(s string) string {
  rtn := []byte{}

  for i := 0; i < len(s); i++ {
    var v int

    if Peek(s, i+2) {
      v, _ = strconv.Atoi(s[i:i+2])
      i += 2
    } else {
      v, _ = strconv.Atoi(s[i:i+1])
    }

    rtn = append(rtn, byte('a' + v - 1))
  }

  return string(rtn)
}

func Peek(s string, i int) bool {
  if i >= len(s) {
    return false
  }
  return s[i] == '#'
}

// T: O(n)
// M: O(1)
func freqAlphabetsBackward(s string) string {
  rtn := []byte{}

  for i := len(s)-1; i >= 0; i-- {
    var v int

    if s[i] == '#' {
      v, _ = strconv.Atoi(s[i-2:i])
      i -= 2
    } else {
      v, _ = strconv.Atoi(s[i:i+1])
    }

    rtn = append([]byte{byte('a'+v-1)}, rtn...)
  }

  return string(rtn)
}

// -- end --

