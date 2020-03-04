package main

// T: O(n)
// M: O(1)
// -- start --

func removeOuterParentheses(S string) string {
  r, c := "", 0

  for i := 0; i < len(S); i++ {
    if S[i] == '(' {
      c++
      continue
    }

    c--
    if c == 0 {
      i, r, S = -1, r + S[1:i], S[i+1:]
    }
  }

  return r
}

// -- end --

