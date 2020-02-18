package main

// T: O(n * l)
// M: O(1)
// -- start --

func isAlienSorted(words []string, order string) bool {
  m := [26]int{}
  for i, v := range order {
    m[v-'a'] = i
  }

  for i := 0; i < len(words)-1; i++ {
    if cmp(m, words[i], words[i+1]) < 0 {
      return false
    }
  }

  return true
}

func cmp(m [26]int, a, b string) int {
  for i := 0; i < len(a) && i < len(b); i++ {
    if m[a[i]-'a'] > m[b[i]-'a'] {
      return -1
    }
    if m[a[i]-'a'] < m[b[i]-'a'] {
      return 1
    }
  }

  if len(a) > len(b) {
    return -1
  }
  if len(a) == len(b) {
    return 0
  }
  return 1
}

// -- end --

