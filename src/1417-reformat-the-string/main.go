package main

// T: O(n)
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func reformat(s string) string {
  bs := []byte{}

  lc, dc := 0, 0
  for i := 0; i < len(s); i++ {
    if isLetter(s[i]) {
      lc++
    } else {
      dc++
    }
  }

  if Abs(lc - dc) > 1 {
    return ""
  }

  di, li := 0, 0
  for i := 0; i < len(s); i++ {
    for di < len(s) && isLetter(s[di]) {
      di++
    }

    for li < len(s) && isDigit(s[li]) {
      li++
    }

    if di >= len(s) && li >= len(s) {
      break
    }

    if di >= len(s) {
      bs = append(bs, s[li])
      break
    }
    if li >= len(s) {
      bs = append(bs, s[di])
      break
    }

    if dc > lc {
      bs = append(bs, s[di], s[li])
    } else {
      bs = append(bs, s[li], s[di])
    }

    di++
    li++
  }

  return string(bs)

}

// T: O(n)
// M: O(n)
func reformatSplit(s string) string {
  digs := []byte{}
  ltrs := []byte{}

  for _, v := range s {
    if isLetter(byte(v)) {
      ltrs = append(ltrs, byte(v))
    } else {
      digs = append(digs, byte(v))
    }
  }

  if Abs(len(digs) - len(ltrs)) > 1 {
    return ""
  }

  rtn := []byte{}

  if len(digs) > len(ltrs) {
    rtn = append(rtn, digs[0])
    digs = digs[1:]
  }

  for i := 0; i < len(digs); i++ {
    rtn = append(rtn, ltrs[i], digs[i])
  }

  if len(digs) < len(ltrs) {
    rtn = append(rtn, ltrs[len(ltrs)-1])
  }

  return string(rtn)
}

func isLetter(v byte) bool {
  if 'a' <= v && v <= 'z' {
    return true
  }
  return false
}

func isDigit(v byte) bool {
  if '0' <= v && v <= '9' {
    return true
  }
  return false
}

func Abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}


// -- end --

