package main
import "testing"
import . "main/pkg/testing_helper"

func TestBulbSwitch(_t *testing.T) {
  i, o := 3, 1
  T(_t, S(i), S(bulbSwitch(i)), S(o))

  i, o = 5, 2
  T(_t, S(i), S(bulbSwitch(i)), S(o))

  i, o = 4, 2
  T(_t, S(i), S(bulbSwitch(i)), S(o))

  i, o = 1, 1
  T(_t, S(i), S(bulbSwitch(i)), S(o))
}

func TestBulbSwitchBruteForce(_t *testing.T) {
  i, o := 3, 1
  T(_t, S(i), S(bulbSwitchBruteForce(i)), S(o))

  i, o = 5, 2
  T(_t, S(i), S(bulbSwitchBruteForce(i)), S(o))

  i, o = 4, 2
  T(_t, S(i), S(bulbSwitchBruteForce(i)), S(o))

  i, o = 1, 1
  T(_t, S(i), S(bulbSwitchBruteForce(i)), S(o))
}

