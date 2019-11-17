package main
import "fmt"

func main() {
  i, o := []string{"gin", "zen", "gig", "msg"}, 2
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", uniqueMorseRepresentations(i))
  fmt.Printf("Expect: %d\n", o)
}

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

