package main
import "testing"
import . "main/pkg/testing_helper"

func TestNumUniqueEmails(t *testing.T) {
  i, o := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}, 2
  T(t, S(i), S(numUniqueEmails(i)), S(o))

  i, o = []string{"test.email+alex@leetcode.com", "test.email@leetcode.com"}, 1
  T(t, S(i), S(numUniqueEmails(i)), S(o))
}

func TestNumUniqueEmailsStrings(t *testing.T) {
  i, o := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}, 2
  T(t, S(i), S(numUniqueEmailsStrings(i)), S(o))

  i, o = []string{"test.email+alex@leetcode.com", "test.email@leetcode.com"}, 1
  T(t, S(i), S(numUniqueEmailsStrings(i)), S(o))
}

