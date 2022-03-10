package main

// T: O(n + w**2) n is the length of S, w is the number of words
// M: O(w)
// -- start --

func toGoatLatin(S string) string {
  S = S + " "

  rtn := ""

  start := 0
  count := 1
  for i, v := range S {
    if v != ' ' {
      continue
    }

    if IsVowel(S[start]) {
      rtn += S[start:i] + "ma"
    } else {
      rtn += S[start+1:i] + S[start:start+1] + "ma"
    }
    start = i+1

    for i := 0; i < count; i++ {
      rtn += "a"
    }

    count++
    rtn += " "
  }

  return rtn[:len(rtn)-1]
}

func IsVowel(v byte) bool {
  switch v {
  case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
    return true
  }
  return false
}

// -- end --

