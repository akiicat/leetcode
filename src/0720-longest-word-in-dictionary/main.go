package main
import "fmt"
import "sort"

func main() {
  i, o := []string{"w","wo","wor","worl", "world"}, "world"
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %s\n", longestWord(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}, "apple"
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %s\n", longestWord(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = []string{"a", "b", "cc"}, "a"
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %s\n", longestWord(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = []string{"w","wo","wox","worl", "world"}, "world"
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %s\n", longestWord(i))
  fmt.Printf("Expect: %s\n", o)
}

// https://leetcode.com/articles/longest-word-in-dictionary/
// T: O(n) 
// M: O(l)
// -- start --

// Trie + Depth-First Search
// T: O(n)
// M: O(l)
type TrieNode struct {
	Children [26]*TrieNode
  idx      int //store the index of the string, instead of the string
	l        int //store the length of the string
}

func (node *TrieNode) Insert(word string, idx int) {
  start := node

  for i := 0; i < len(word); i++ {
    node := start.Children[word[i]-'a']

    // found node
    if node != nil {
      start = node
      continue
    }

    // not found
    for i < len(word) {
      next := &TrieNode{}
      start.Children[word[i]-'a'] = next
      start = next
      i++
    }
  }

  // word ended in this node
  start.l = len(word)
  start.idx = idx
}

func (root *TrieNode) Search() *TrieNode {
  if root.l == 0 {
    return root
  }
  cur := root
  for _, node := range root.Children {
    if node == nil {
      continue
    }

    res := node.Search()
    if res.l > cur.l {
      cur = res
    }
  }
  return cur
}

func longestWord(words []string) string {
  root := &TrieNode{}
  for i, w := range words {
    root.Insert(w, i)
  }
  root.l = -1
  res := root.Search()
  if res.l <= 0 {
    return ""
  }
  return words[res.idx]
}

// T: O(n*l) 
// M: O(n*l)
func longestWordSort(words []string) string {
  sort.Strings(words)

  m := make(map[string]bool)

  for _, v := range words {
    m[v] = true
  }

  rtn := ""
  for _, v := range words {
    if len(v) == 1 || m[v[:len(v)-1]] {
      if len(v) > len(rtn) || len(rtn) == len(v) && rtn > v {
        rtn = v
      }
    } else {
      m[v] = false
    }
  }

  return rtn
}

// -- end --

