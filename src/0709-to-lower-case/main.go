package main

// T: O(n)
// M: O(1)
// -- start --

func toLowerCase(str string) string {
  rtn := []rune{}
  for _, v := range str {
    if 'A' <= v && v <= 'Z' {
      v = v - 'A' + 'a'
    }
    rtn = append(rtn, v)
  }
  return string(rtn)
}

// -- end --

