package main

// T: O(n) n for all of chars
// M: O(1)
// -- start --

func commonChars(A []string) []string {
  m := make([]int, 26)

  for i := 0; i < 26; i++ {
    m[i] = (1 << 31) - 1
  }

  for _, str := range A {
    tmp := make([]int, 26)
    for _, c := range []byte(str) {
      tmp[c - 'a']++
    }

    for k, num := range tmp {
      m[k] = min(num, m[k])
    }
  }

  rtn := []string{}

  for i := 0; i < 26; i++ {
    if m[i] == (1 << 31) - 1 {
      continue
    }

    for j := 0; j < m[i]; j++ {
      rtn = append(rtn, string(byte(i+'a')))
    }
  }

  return rtn
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

