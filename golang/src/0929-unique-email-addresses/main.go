package main
import "strings"

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
  ignored := false

  for i, b := range email {
    if b == '@' {
      s = append(s, []byte(email[i:])...)
      break
    }
    if ignored || b == '.' {
      continue
    }
    if b == '+' {
      ignored = true
      continue
    }
    s = append(s, byte(b))
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

