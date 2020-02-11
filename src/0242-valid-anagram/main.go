package main

// T: O(n)
// M: O(1) for 26
// -- start --

func isAnagram(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }

  m := [26]int{}

  for _, v := range s {
    m[v-'a']++
  }

  for _, v := range t {
    m[v-'a']--
  }

  for _, v := range m {
    if v != 0 {
      return false
    }
  }

  return true
}

// -- end --

