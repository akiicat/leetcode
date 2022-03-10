package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsOneBitCharacter(_t *testing.T) {
  i, o := []int{1,0,0}, true
  T(_t, S(i), S(isOneBitCharacter(i)), S(o))

  i, o = []int{1,1,1,0}, false
  T(_t, S(i), S(isOneBitCharacter(i)), S(o))

  i, o = []int{0}, true
  T(_t, S(i), S(isOneBitCharacter(i)), S(o))
}

func TestIsOneBitCharacterGreedy(_t *testing.T) {
  i, o := []int{1,0,0}, true
  T(_t, S(i), S(isOneBitCharacterGreedy(i)), S(o))

  i, o = []int{1,1,1,0}, false
  T(_t, S(i), S(isOneBitCharacterGreedy(i)), S(o))

  i, o = []int{0}, true
  T(_t, S(i), S(isOneBitCharacterGreedy(i)), S(o))
}

func TestIsOneBitCharacterOneBool(_t *testing.T) {
  i, o := []int{1,0,0}, true
  T(_t, S(i), S(isOneBitCharacterOneBool(i)), S(o))

  i, o = []int{1,1,1,0}, false
  T(_t, S(i), S(isOneBitCharacterOneBool(i)), S(o))

  i, o = []int{0}, true
  T(_t, S(i), S(isOneBitCharacterOneBool(i)), S(o))
}

func TestIsOneBitCharacterTwoBool(_t *testing.T) {
  i, o := []int{1,0,0}, true
  T(_t, S(i), S(isOneBitCharacterTwoBool(i)), S(o))

  i, o = []int{1,1,1,0}, false
  T(_t, S(i), S(isOneBitCharacterTwoBool(i)), S(o))

  i, o = []int{0}, true
  T(_t, S(i), S(isOneBitCharacterTwoBool(i)), S(o))
}

