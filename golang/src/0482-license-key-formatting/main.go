package main

// T: O(n)
// M: O(n)
// -- start --

func licenseKeyFormatting(S string, K int) string {
  c := 0
  s := []byte{} // stack

  for i := len(S)-1; i >= 0; i-- {
    if S[i] == '-' {
      continue
    }

    if c == K {
      s, c = append(s, '-'), 0
    }

    if S[i] >= 97 {
      s, c = append(s, S[i]-32), c + 1
    } else {
      s, c = append(s, S[i]), c + 1
    }
  }

  // reverse stack
  l, r := 0, len(s)-1
  for l < r {
    s[l], s[r] = s[r], s[l]
    l++
    r--
  }

  return string(s)
}

// -- end --

