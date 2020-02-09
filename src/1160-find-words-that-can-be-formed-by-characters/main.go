package main

// T: O(n * m) n is the length of words, m is the size of word
// M: O(1) 26
// -- start --

func countCharacters(words []string, chars string) int {
  m := [26]int{}

  for _, v := range chars {
    m[v-'a']++
  }

  sum := 0
  for _, word := range words {
    sum += valid(m, word)
  }
  return sum
}

func valid(m [26]int, word string) int {
  t := [26]int{}
  for _, v := range word {
    t[v-'a']++
    if t[v-'a'] > m[v-'a'] {
      return 0
    }
  }
  return len(word)
}

// -- end --

