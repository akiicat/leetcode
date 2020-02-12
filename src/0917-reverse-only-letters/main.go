package main

// T: O(n)
// M: O(1)
// -- start --

func reverseOnlyLetters(S string) string {
  m := []byte(S)

  l, r := 0, len(S)-1
  for l < r {
    if !isAlphabet(m[l]) {
      l++
      continue
    }
    if !isAlphabet(m[r]) {
      r--
      continue
    }
    m[l], m[r] = m[r], m[l]
    l++
    r--
  }

  return string(m)
}

func isAlphabet(a byte) bool {
  return ('a' <= a && a <= 'z') || ('A' <= a && a <= 'Z')
}

// -- end --

