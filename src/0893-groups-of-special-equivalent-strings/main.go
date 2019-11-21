package main
import "fmt"

func main() {
  i, o := []string{"abcd","cdab","cbad","xyzz","zzxy","zzyx"}, 3
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", numSpecialEquivGroups(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []string{"abc","acb","bac","bca","cab","cba"}, 3
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", numSpecialEquivGroups(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(n)
// -- start --

func numSpecialEquivGroups(A []string) int {
  m := make(map[[52]int]bool)

  for _, word := range A {
    key := [52]int{}
    for i, v := range word {
      key[int(v - 'a') + 26 * (i % 2)] += 1
    }
    m[key] = true
  }

  return len(m)
}

// T: O(n)
// M: O(n)
func numSpecialEquivGroupsStringMap(A []string) int {
  m := make(map[string]bool)

  for _, word := range A {
    a, b := convert(word)
    m[a + " " + b] = true
  }

  return len(m)
}

func convert(word string) (string, string) {
  m := [52]int{}

  for i, w := range word {
    m[int(w - 'a') + 26 * (i % 2)]++
  }

  even, odd := []byte{}, []byte{}

  for i, w := range m {
    if w == 0 {
      continue
    }
    if i < 26 {
      for w >= 0 {
        even = append(even, byte('a' + i % 26))
        w--
      }
    } else {
      for w >= 0 {
        odd = append(odd, byte('a' + i % 26))
        w--
      }
    }
  }

  return string(even), string(odd)
}

// -- end --

