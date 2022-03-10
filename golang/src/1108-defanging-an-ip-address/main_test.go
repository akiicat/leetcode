package main
import "testing"
import . "main/pkg/testing_helper"

func TestDefangIPaddr(t *testing.T) {
  i, o := "1.1.1.1", "1[.]1[.]1[.]1"
  T(t, S(i), S(defangIPaddr(i)), S(o))

  i, o = "255.100.50.0", "255[.]100[.]50[.]0"
  T(t, S(i), S(defangIPaddr(i)), S(o))
}

func TestDefangIPaddrConcat(t *testing.T) {
  i, o := "1.1.1.1", "1[.]1[.]1[.]1"
  T(t, S(i), S(defangIPaddrConcat(i)), S(o))

  i, o = "255.100.50.0", "255[.]100[.]50[.]0"
  T(t, S(i), S(defangIPaddrConcat(i)), S(o))
}

func TestDefangIPaddrBuilder(t *testing.T) {
  i, o := "1.1.1.1", "1[.]1[.]1[.]1"
  T(t, S(i), S(defangIPaddrBuilder(i)), S(o))

  i, o = "255.100.50.0", "255[.]100[.]50[.]0"
  T(t, S(i), S(defangIPaddrBuilder(i)), S(o))
}

func TestDefangIPaddrBuffer(t *testing.T) {
  i, o := "1.1.1.1", "1[.]1[.]1[.]1"
  T(t, S(i), S(defangIPaddrBuffer(i)), S(o))

  i, o = "255.100.50.0", "255[.]100[.]50[.]0"
  T(t, S(i), S(defangIPaddrBuffer(i)), S(o))
}

func TestDefangIPaddrStrings(t *testing.T) {
  i, o := "1.1.1.1", "1[.]1[.]1[.]1"
  T(t, S(i), S(defangIPaddrStrings(i)), S(o))

  i, o = "255.100.50.0", "255[.]100[.]50[.]0"
  T(t, S(i), S(defangIPaddrStrings(i)), S(o))
}

