package main

// T: O(n) n is total chars
// M: O(1)
// -- start --

func findWords(words []string) []string {
  rtn := []string{}

  for _, word := range words {
    if isValid(word) {
      rtn = append(rtn, word)
    }
  }

  return rtn
}

func isValid(word string) bool {
  m := []int8{1,0,0,1,2,1,1,1,2,1,1,1,0,0,2,2,2,2,1,2,2,0,2,0,2,0}

  cur := int8(-1)

  for i, c := range word {

    if c <= 'Z' {
      c += 32
    }
    c -= 'a'

    if i == 0 {
      cur = m[c]
      continue
    }

    if cur != m[c] {
      return false
    }
  }

  return true
}

// -- end --

