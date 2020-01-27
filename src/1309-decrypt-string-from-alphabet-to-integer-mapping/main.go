package main
import "fmt"
import "strconv"

func main() {
  i, o := "10#11#12", "jkab"
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %s\n", freqAlphabets(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = "1326#", "acz"
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %s\n", freqAlphabets(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = "25#", "y"
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %s\n", freqAlphabets(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#", "abcdefghijklmnopqrstuvwxyz"
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %s\n", freqAlphabets(i))
  fmt.Printf("Expect: %s\n", o)
}

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

