package main

// T: O(n) n is the length of string
// M: O(1)
// -- start --

// implement len(strings.Split(s, " "))
func countSegments(s string) int {
  c := 0

  for i, v := range s {
    if (i == 0 || s[i-1] == ' ') && v != ' ' {
      c++
    }
  }

  return c
}

func countSegmentsCountWords(s string) int {
  w, c := 0, 0

  for _, v := range s {
    if v == ' ' {
      if w != 0 {
        c++
      }
      w = 0
    } else {
      w++
    }
  }

  if w != 0 {
    c++
  }

  return c
}

// -- end --

