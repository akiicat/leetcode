package main

// T: O(n * l)
// M: O(n * l)
// -- start --

func uniqueMorseRepresentations(words []string) int {
  m := make(map[string]bool)
  for _, word := range words {
    m[convert(word)] = true
  }
  return len(m)
}

func convert(word string) string {
  a := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
  s := ""

  for _, v := range word {
    s = s + a[v - 'a']
  }

  return s
}

// -- end --

