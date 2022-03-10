package main
import "testing"
import . "main/pkg/testing_helper"

func TestGuessNumber(_t *testing.T) {
  i, o := 10, 6
  setGuess(o)
  T(_t, S(i), S(guessNumber(i)), S(o))

  i, o = 1, 1
  setGuess(o)
  T(_t, S(i), S(guessNumber(i)), S(o))
}

