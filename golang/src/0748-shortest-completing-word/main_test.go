package main
import "testing"
import . "main/pkg/testing_helper"

func TestShortestCompletingWord(_t *testing.T) {
  i, licensePlate, o := []string{"step", "steps", "stripe", "stepple"}, "1s3 PSt", "steps"
  T(_t, S(licensePlate, i), S(shortestCompletingWord(licensePlate, i)), S(o))

  i, licensePlate, o = []string{"looks", "pest", "stew", "show"}, "1s3 456", "pest"
  T(_t, S(licensePlate, i), S(shortestCompletingWord(licensePlate, i)), S(o))
}

func TestShortestCompletingWordMap(_t *testing.T) {
  i, licensePlate, o := []string{"step", "steps", "stripe", "stepple"}, "1s3 PSt", "steps"
  T(_t, S(licensePlate, i), S(shortestCompletingWordMap(licensePlate, i)), S(o))

  i, licensePlate, o = []string{"looks", "pest", "stew", "show"}, "1s3 456", "pest"
  T(_t, S(licensePlate, i), S(shortestCompletingWordMap(licensePlate, i)), S(o))
}

