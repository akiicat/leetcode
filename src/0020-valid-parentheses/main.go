package main
import "fmt"

func main() {
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: true\n", "()", isValid("()"))
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: true\n", "()[]{}", isValid("()[]{}"))
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: true\n", "{[]}", isValid("{[]}"))
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: false\n", "[", isValid("["))
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: false\n", "]", isValid("]"))
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: false\n", "(]", isValid("(]"))
  fmt.Printf("Input:  %s\nOutput: %t\nExpect: false\n", "([)]", isValid("([)]"))
}

// T: O(N) because we simply traverse the given string one character at a time and push and pop operations on a stack take O(1) time.
// M: O(N) as we push all opening brackets onto the stack and in the worst case, we will end up pushing all the brackets onto the stack. e.g. ((((((((((.
// -- start --

func isValid(s string) bool {
  bs := []byte{}

  for _, char := range []byte(s) {
    if char == '(' || char == '{' || char == '[' {
      bs = append(bs, char)
    } else if char == ')' || char =='}' || char == ']' {
      if len(bs) == 0 || !valid(bs[len(bs)-1], char) {
        return false
      }
      bs = bs[:len(bs)-1]
    }
  }

  return len(bs) == 0
}

func valid(c1, c2 byte) bool {
  return c1 == '{' && c2 == '}' || c1 == '(' && c2 == ')' || c1 == '[' && c2 == ']'
}


// -- end --

