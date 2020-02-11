package main
import "testing"
import . "main/pkg/testing_helper"

func TestStrStr(_t *testing.T) {
  h, n, o := "hello", "ll", 2
  T(_t, S(h, n), S(strStr(h, n)), S(o))

  h, n, o = "aaaaa", "bba", -1
  T(_t, S(h, n), S(strStr(h, n)), S(o))

  h, n, o = "", "", 0
  T(_t, S(h, n), S(strStr(h, n)), S(o))

  h, n, o = "", "a", -1
  T(_t, S(h, n), S(strStr(h, n)), S(o))

  h, n, o = "a", "a", 0
  T(_t, S(h, n), S(strStr(h, n)), S(o))

  h, n, o = "abababcabaa", "ababc", 2
  T(_t, S(h, n), S(strStr(h, n)), S(o))
}

func TestStrStrKmpWithPrefixTable(_t *testing.T) {
  h, n, o := "hello", "ll", 2
  T(_t, S(h, n), S(strStrKmpWithPrefixTable(h, n)), S(o))

  h, n, o = "aaaaa", "bba", -1
  T(_t, S(h, n), S(strStrKmpWithPrefixTable(h, n)), S(o))

  h, n, o = "", "", 0
  T(_t, S(h, n), S(strStrKmpWithPrefixTable(h, n)), S(o))

  h, n, o = "", "a", -1
  T(_t, S(h, n), S(strStrKmpWithPrefixTable(h, n)), S(o))

  h, n, o = "a", "a", 0
  T(_t, S(h, n), S(strStrKmpWithPrefixTable(h, n)), S(o))

  h, n, o = "abababcabaa", "ababc", 2
  T(_t, S(h, n), S(strStrKmpWithPrefixTable(h, n)), S(o))
}

func TestStrStrBruteForce(_t *testing.T) {
  h, n, o := "hello", "ll", 2
  T(_t, S(h, n), S(strStrBruteForce(h, n)), S(o))

  h, n, o = "aaaaa", "bba", -1
  T(_t, S(h, n), S(strStrBruteForce(h, n)), S(o))

  h, n, o = "", "", 0
  T(_t, S(h, n), S(strStrBruteForce(h, n)), S(o))

  h, n, o = "", "a", -1
  T(_t, S(h, n), S(strStrBruteForce(h, n)), S(o))

  h, n, o = "a", "a", 0
  T(_t, S(h, n), S(strStrBruteForce(h, n)), S(o))

  h, n, o = "abababcabaa", "ababc", 2
  T(_t, S(h, n), S(strStrBruteForce(h, n)), S(o))
}

func TestStrStrLibrary(_t *testing.T) {
  h, n, o := "hello", "ll", 2
  T(_t, S(h, n), S(strStrLibrary(h, n)), S(o))

  h, n, o = "aaaaa", "bba", -1
  T(_t, S(h, n), S(strStrLibrary(h, n)), S(o))

  h, n, o = "", "", 0
  T(_t, S(h, n), S(strStrLibrary(h, n)), S(o))

  h, n, o = "", "a", -1
  T(_t, S(h, n), S(strStrLibrary(h, n)), S(o))

  h, n, o = "a", "a", 0
  T(_t, S(h, n), S(strStrLibrary(h, n)), S(o))

  h, n, o = "abababcabaa", "ababc", 2
  T(_t, S(h, n), S(strStrLibrary(h, n)), S(o))
}

