package main
import "fmt"

func main() {
  i1, i2, o := "0", "0", "0"
  fmt.Printf("Input:  %s, %s\n", i1, i2)
  fmt.Printf("Output: %s\n", addStrings(i1, i2))
  fmt.Printf("Expect: %s\n", o)

  i1, i2, o = "10", "90", "100"
  fmt.Printf("Input:  %s, %s\n", i1, i2)
  fmt.Printf("Output: %s\n", addStrings(i1, i2))
  fmt.Printf("Expect: %s\n", o)

  i1, i2, o = "9", "99", "108"
  fmt.Printf("Input:  %s, %s\n", i1, i2)
  fmt.Printf("Output: %s\n", addStrings(i1, i2))
  fmt.Printf("Expect: %s\n", o)

  i1, i2, o = "5432", "54", "5486"
  fmt.Printf("Input:  %s, %s\n", i1, i2)
  fmt.Printf("Output: %s\n", addStrings(i1, i2))
  fmt.Printf("Expect: %s\n", o)
}

// T: O(n) n is the max of characters length in all strings.
// M: O(n)
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

