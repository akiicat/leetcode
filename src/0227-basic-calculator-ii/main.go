package main
import "strconv"
import "unicode"
import "bytes"

// T: O(n)
// M: O(1)
// -- start --

func calculate(s string) int {

  var buf bytes.Buffer
  for _, r := range s {
    if r != ' ' {
      buf.WriteRune(r)
    }
  }

  if buf.Len() == 0 {
    return 0
  }

  s  = buf.String()

  i := 0
  cur := GetNum(s, &i)

  res := 0
  for i < len(s) {

    oper := s[i]
    i++

    next := GetNum(s, &i)

    switch oper {
    case '+':
      res += cur
      cur = next
    case '-':
      res += cur
      cur = -next
    case '*':
      cur *= next
    case '/':
      cur /= next
    }
  }

  return res + cur
}

func GetNum(s string, i *int) int {
  start := *i
  for *i < len(s) && unicode.IsNumber(rune(s[*i])) {
    *i++
  }
  cur, _ := strconv.Atoi(s[start:*i])
  return cur
}

// -- end --

