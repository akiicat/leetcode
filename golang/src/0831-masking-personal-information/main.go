package main

// T: O(1)
// M: O(1)
// -- start --

func maskPII(S string) string {
  phone := []byte{}
  for i, v := range S {
    if v == '@' {
      // mask mail
      S = Lower(S)
      name := S[:i]
      name = string([]byte{name[0], '*', '*', '*', '*', '*', name[len(name)-1]})
      return name + S[i:]
    }

    if '0' <= v && v <= '9' {
      phone = append(phone, byte(v))
    }
  }

  // mask mail
  res := "***-***-" + string(phone[len(phone)-4:])
  if len(phone) == 10 {
    return res
  }

  res = "-" + res
  for i := 10; i < len(phone); i++ {
    res = "*" + res
  }
  res = "+" + res
  return res
}

func Lower(s string) string {
  bs := []byte(s)
  for i, c := range bs {
    if 'A' <= c && c <= 'Z' {
      bs[i] = c - 'A' + 'a'
    }
  }
  return string(bs)
}

// -- end --

