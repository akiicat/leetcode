package main

// T: O(n)
// M: O(1)
// -- start --

func myAtoi(str string) int {
  for i, v := range str {
    if !isSpace(v) {
      str = str[i:]
      break
    }
  }

  if len(str) == 0 {
    return 0
  }

  negative := false
  if isSign(rune(str[0])) {
    if str[0] == '-' {
      negative = true
    }
    str = str[1:]
  }

  rtn := 0
  for _, v := range str {
    if !isDigit(v) {
      break
    }
    rtn = 10 * rtn + int(v - '0')

    if !negative && rtn > 1<<31-1 {
      return 1<<31-1
    }
    if negative && rtn > 1<<31 {
      return -1<<31
    }
  }

  if negative {
    return -rtn
  }

  return rtn
}

func isDigit(v rune) bool {
  return '0' <= v && v <= '9'
}

func isSpace(v rune) bool {
  return v == ' '
}

func isSign(v rune) bool {
  return v == '-' || v == '+'
}

func isValid(v rune) bool {
  return isDigit(v) || isSign(v)
}

// -- end --

