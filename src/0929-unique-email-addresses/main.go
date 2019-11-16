package main
import "fmt"
import "strings"

func main() {
  i, o := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}, 2
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", numUniqueEmails(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []string{"test.email+alex@leetcode.com", "test.email@leetcode.com"}, 1
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", numUniqueEmails(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n * l)
// M: O(n * l)
// -- start --

func numUniqueEmails(emails []string) int {
  m := make(map[string]bool)
  for _, email := range emails {
    m[uniq(email)] = true
  }
  return len(m)
}

func uniq(email string) string {
  s := []byte{}
  i := 0
  t := true

  for i < len(email) {
    b := email[i]
    if b == '@' {
      s = append(s, []byte(email[i:])...)
      break
    }
    if b == '+' {
      t = false
    }
    if b == '.' {
      i++
      continue
    }
    if t {
      s = append(s, b)
    }
    i++
  }
  return string(s)
}

func numUniqueEmailsStrings(emails []string) int {
  m := make(map[string]bool)
  for _, email := range emails {
    nameDomain := strings.Split(email, "@")
    realName := strings.Split(nameDomain[0],"+")
    nameNoDot := strings.Join(strings.Split(realName[0],"."),"")

    m[nameNoDot + "@" + nameDomain[1]] = true
  }
  return len(m)
}

// -- end --

