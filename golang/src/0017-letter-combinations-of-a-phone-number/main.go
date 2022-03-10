package main

// T: O(4 ^ n)
// M: O(4 ^ n)
// -- start --

var m = map[byte]string{
  '2': "abc",
  '3': "def",
  '4': "ghi",
  '5': "jkl",
  '6': "mno",
  '7': "pqrs",
  '8': "tuv",
  '9': "wxyz",
}

func letterCombinations(digits string) []string {
  if len(digits) == 0 {
    return nil
  }

  return R([]string{""}, digits)
}

func R(a []string, digits string) []string {
  if len(digits) == 0 {
    return a
  }

  s := []string{}

  for _, c := range m[digits[0]] {
    for _, v := range a {
      s = append(s, v + string([]rune{c}))
    }
  }

  return R(s, digits[1:])
}

// -- end --

