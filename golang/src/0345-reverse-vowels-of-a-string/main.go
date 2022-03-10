package main
import "bytes"

// T: O(n)
// M: O(1)
// -- start --

func reverseVowels(s string) string {
  l, r := 0, len(s) - 1

  rtn := []byte(s)

  for l < r {

    if !bytes.ContainsRune([]byte("AEIOUaeiou"), rune(rtn[l])) {
      l++
      continue
    }

    if !bytes.ContainsRune([]byte("AEIOUaeiou"), rune(rtn[r])) {
      r--
      continue
    }

    rtn[l], rtn[r] = rtn[r], rtn[l]

    l++
    r--
  }

  return string(rtn)
}

// -- end --
