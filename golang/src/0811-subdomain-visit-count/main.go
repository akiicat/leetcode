package main
import "strings"
import "strconv"

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

