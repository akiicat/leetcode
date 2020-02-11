package main
import "testing"
import . "main/pkg/testing_helper"

func TestSubdomainVisits(_t *testing.T) {
  i, o := []string{"9001 discuss.leetcode.com"}, []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"}
  T(_t, S(i), S(SortStr(subdomainVisits(i))), S(SortStr(o)))

  i, o = []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}, []string{"901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"}
  T(_t, S(i), S(SortStr(subdomainVisits(i))), S(SortStr(o)))
}

