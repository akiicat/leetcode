package main
import "math"
import "unicode"

// T: O(n * 2^n)
// M: O(1)
// -- start --

// T: O(n * 2^n)
// M: O(1)
func letterCasePermutation(S string) []string {
  m := []string{}
  R(S, &m, 0)
  return m
}

// DFS
func R(S string, m *[]string, i int) {
  if i == len(S) {
    *m = append(*m, S)
    return
  }

  if !unicode.IsLetter(rune(S[i])) {
    R(S, m, i+1)
    return
  }

  t := convertCharByte(S[i])
  s := S[:i] + string(t) + S[i+1:]

  R(s, m, i+1)
  R(S, m, i+1)
}

func convertCharByte(v byte) byte {
  if 'a' <= v && v <= 'z' {
    return v - 'a' + 'A'
  }
  return v - 'A' + 'a'
}

// T: O(n * 2^n)
// M: O(1)
func letterCasePermutationBit(S string) []string {
  rtn := []string{}
  m := []int{}

  n := 0
  for i, v := range S {
    if isChar(v) {
      n++
      m = append(m, i)
    }
  }

  n = int(math.Pow(2, float64(n)))

  for i := 0; i < n; i++ {
    s := []rune{}

    j := i
    for _, v := range S {
      if !isChar(v) {
        s = append(s, v)
        continue
      }

      if j & 0x1 == 1 {
        v = convertChar(v)
      }
      j >>= 1

      s = append(s, v)
    }

    rtn = append(rtn, string(s))
  }

  return rtn
}

func isChar(v rune) bool {
  if 'a' <= v && v <= 'z' || 'A' <= v && v <= 'Z' {
    return true
  }
  return false
}

func convertChar(v rune) rune {
  if 'a' <= v && v <= 'z' {
    return v - 'a' + 'A'
  }

  return v - 'A' + 'a'
}

// -- end --

