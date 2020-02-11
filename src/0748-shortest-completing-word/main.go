package main

// T: O(n)
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func shortestCompletingWord(licensePlate string, words []string) string {
  m := Count(licensePlate)

  rtn, min := "", 16
  for _, word := range words {
    if len(word) >= min {
      continue
    }

    t := Count(word)
    if !Vaild(m, t) {
      continue
    }

    rtn, min = word, len(word)
  }

  return rtn
}

func Count(s string) [26]int8 {
  m := [26]int8{}
  for _, v := range s {
    if 'A' <= v && v <= 'Z' {
      m[int(v-'A')]++
    } else if 'a' <= v && v <= 'z' {
      m[int(v-'a')]++
    }
  }
  return m
}

func Vaild(m, t [26]int8) bool {
  for k, _ := range m {
    if m[k] > t[k] {
      return false
    }
  }
  return true
}

// T: O(n)
// M: O(1)
func shortestCompletingWordMap(licensePlate string, words []string) string {
  m := CountMap(licensePlate)

  rtn, min := "", 16
  for _, word := range words {
    if len(word) >= min {
      continue
    }

    t := CountMap(word)
    if !VaildMap(m, t) {
      continue
    }

    rtn, min = word, len(word)
  }

  return rtn
}

func CountMap(s string) map[rune]int {
  m := make(map[rune]int)
  for _, v := range s {
    if 'A' <= v && v <= 'Z' {
      m[v-'A']++
    } else if 'a' <= v && v <= 'z' {
      m[v-'a']++
    }
  }
  return m
}

func VaildMap(m, t map[rune]int) bool {
  for k, _ := range m {
    if m[k] > t[k] {
      return false
    }
  }
  return true
}


// -- end --

