package main

// T: O(n + m)
// M: O(1)
// -- start --

func gcdOfStrings(str1 string, str2 string) string {
  for len(str1) != len(str2) {
    if len(str1) < len(str2) {
      str1, str2 = str2, str1
    }

    str1 = str1[len(str2):]
  }

  if str1[:len(str2)] != str2 {
    return ""
  }

  return str1
}

// -- end --
