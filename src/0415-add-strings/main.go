package main
import "fmt"

func main() {
  s := []string{"0", "0"}
  fmt.Printf("Input:  %v\nOutput: %s\nExpect: 0\n", s, addStrings(s[0], s[1]))

  s = []string{"10", "90"}
  fmt.Printf("Input:  %v\nOutput: %s\nExpect: 100\n", s, addStrings(s[0], s[1]))

  s = []string{"9", "99"}
  fmt.Printf("Input:  %v\nOutput: %s\nExpect: 108\n", s, addStrings(s[0], s[1]))

  s = []string{"5432", "54"}
  fmt.Printf("Input:  %v\nOutput: %s\nExpect: 5486\n", s, addStrings(s[0], s[1]))
}

// T: O(N) N is the max of characters length in all strings.
// M: O(N)
// -- start --

func addStrings(num1 string, num2 string) string {
  l, r := len(num1), len(num2)
  rtn := make([]byte, 0, max(l, r) + 1)

  var carry byte
  for l, r = l-1, r-1; l >= 0 || r >= 0 || carry == 1; l, r = l-1, r-1 {
    var a, b byte

    if l >= 0 {
      a = num1[l] - '0'
    }

    if r >= 0 {
      b = num2[r] - '0'
    }

    sum := a + b + carry
    rtn = append(rtn, sum % 10 + '0')
    carry = sum / 10
  }

  for i, j := 0, len(rtn)-1; i < j; i, j = i+1, j-1 {
    rtn[i], rtn[j] = rtn[j], rtn[i]
  }

  return string(rtn)
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

