package main
import "fmt"
import "strings"
import "strconv"

func main() {
  i, o := []string{"9001 discuss.leetcode.com"}, []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", subdomainVisits(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}, []string{"901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", subdomainVisits(i))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n)
// M: O(n)
// -- start --

func subdomainVisits(cpdomains []string) []string {
  m := make(map[string]int)

  for _, domain := range cpdomains {
    d := strings.Index(domain, " ")
    num, _ := strconv.Atoi(domain[:d])
    domain = "." + domain[d+1:]

    for i := len(domain)-1; i >= 0; i-- {
      if domain[i] == '.' {
        m[domain[i+1:]] += num
      }
    }
  }

  rtn := []string{}
  for k, v := range m {
    rtn = append(rtn, strconv.Itoa(v) + " " + k)
  }

  return rtn
}

// -- end --

